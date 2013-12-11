// Copyright 2013 The gohaxelib Authors listed in the AUTHORS file. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be found in the LICENSE file.
  
import haxe.rtti.XmlParser;
import haxe.rtti.CType;
import sys.io.File;
import sys.io.FileOutput;
import sys.FileSystem;

class Haxelib2go {
	
    static	var parser = new haxe.rtti.XmlParser(); // this variable is reinitialised for each file
	
    static function main() { // Main entry point
    	var out:String="";
		
		/* The standard library */
        out+=xmlFile("api-xml/cpp.xml","cpp","P");
        out+=xmlFile("api-xml/cross.xml","cross","X");
        out+=xmlFile("api-xml/cs.xml","cs","R");
        out+=xmlFile("api-xml/flash9.xml","flash9","F");
        out+=xmlFile("api-xml/java.xml","java","J");
        out+=xmlFile("api-xml/js.xml","js","S");
        out+=xmlFile("api-xml/neko.xml","neko","N");
        out+=xmlFile("api-xml/php.xml","php","H");
        FileUtils.writeTextFileWithBOM("package _haxeapi\n"+out,"_haxeapi/def.go");
        
		/* OpenFL */
		out=xmlFile("api-xml/openfl.xml","","X");
        FileUtils.writeTextFileWithBOM("package _openfl\n"+out,"_openfl/def.go");
    }
	
    static function xmlFile(f:String,t:String,h:String):String { // process a single file
        
		// standard code to load in the XML data and point at it's root node
		parser = new haxe.rtti.XmlParser();
		var data = sys.io.File.getContent(f);
        var doc = Xml.parse(data).firstElement();
        parser.process(doc,t);
        var tr = parser.root;
		
		// itterate through the TypeTree entries, appending to the output string as we go 
   		var out:String = "";
        for(t in tr)	
	       out += tt(t,h);
	    return out;    	
    }
	
    static function tt(t:haxe.rtti.TypeTree,h:String):String { // process an individual TypeTree 
    	return
    		switch(t){
    		case TTypedecl(tt):
    			if(tt.isPrivate) 
  					"";
				else
					"type "+h+ d2_(tt.path) + " " + haxe2goType(tt.type,h) + " // TTypedecl\n";
			case TPackage(name,full,subs):
				var ret:String = "";
				for(s in subs)
					ret += tt(s,h);
				ret;
			case TEnumdecl(e):
				if(e.isPrivate)
					"";
				else
					"type "+h+ d2_(e.path) + " uintptr // TEnumdecl\n";
			case TClassdecl(c):
				if(!c.isPrivate){
					var ret:String = "";
					for(s in c.statics) {
						if(s.isPublic){
							switch(s.type){
							case CFunction(args,_ret):
								ret += "func "+h+d2_(c.path)+"_"+d2_(s.name)+ "(";
								var hadFirst:Bool=false;
								for(a in args) {
									if(a.name!=""){
										if(hadFirst)
											ret+=",";
										ret += ""+h+a.name + " " + haxe2goType(a.t,h);
										hadFirst=true;
									}	
								}
								if(args.length>0) 
									ret += "";
								ret += ")";
								ret+=haxe2goFuncBody(_ret,h);
								ret += "\n";
							default:
								// add getters & setters for static vars
								ret += "func "+h+ d2_(c.path) + "_"+d2_(s.name) +"_goget() ";
								ret += haxe2goFuncBody(s.type,h);
								ret += "\n";
								ret += "func "+h+ d2_(c.path) + "_"+d2_(s.name) +"_goset(";
								ret += haxe2goType(s.type,h)+"){}";
								ret += "\n";
							}
						}
					}
					for(postfix in [""].concat(c.params)){
						var ul = postfix==""?"":"_";
							ret += "type "+h+ d2_(c.path) + ul + postfix + " uintptr // TClassdecl\n";
					}
					for(m in c.fields)
						if(m.isPublic)	{					
							switch(m.type){
							case CFunction(args,_ret):
								if(m.name=="new")
									ret += "func "+h+d2_(c.path)+"_" +d2_(m.name) +"(";
								else
									ret += "func (x "+h+ d2_(c.path) + ") "+h+d2_(m.name) +"(";
								var hadFirst:Bool=false;
								for(a in args) {
									if(a.name!=""){
										if(hadFirst)
											ret+=",";
										ret += h+a.name + " " + haxe2goType(a.t,h);
										hadFirst=true;
									}	
								}
								if(args.length>0) 
									ret += "";
								ret += ")";
								if(m.name=="new")
									ret+=" "+h+d2_(c.path)+" { return "+h+d2_(c.path)+"(0); }";
								else
									ret+=haxe2goFuncBody(_ret,h);
								ret += "\n";
							default:
								ret += "func (x "+h+ d2_(c.path) + ") "+h+d2_(m.name) +"_goget() ";
								ret += haxe2goFuncBody(m.type,h);
								ret += "\n";
								ret += "func (x "+h+ d2_(c.path) + ") "+h+d2_(m.name) +"_goset(";
								ret += haxe2goType(m.type,h)+"){}";
								ret += "\n";
							}
						}
					ret;
				} else "";
			case TAbstractdecl(a): 
    			if(a.isPrivate)
    				return ""
    			else
	    			"type "+h+ d2_(a.path) + " uintptr // TAbstractdecl\n";
    		}
    }
	
    static function d2_(x:String):String { // convert dots to underlines, underline to tripple underline
    	if(x==null)
    		return null;
    	if(x.length>2)
	    	return x.split("_").join("...").split(".").join("_"); // transform orig "_" to "___" 
    	else
    		return x;
    }
	
    static function haxe2goType(t:haxe.rtti.CType,h:String):String { // mapping of haxe to Go types
    	// TODO work out which of "Void" or Int/Float/Bool is actually the one used...
		switch(t){
		case CUnknown:
		case CTypedef(name,params):
			switch(name) {
				case "Void": return "";
			}
		case CFunction(args,_ret):
			return "interface{}";
		case CEnum(name,params):
			switch(name) {
				case "Void": return "";
				case "Bool": return "bool";
				default: return className(name,params,h); 
			}
		case CDynamic(t):
			return "interface{}";
		case CClass(name,params):
			switch(name){
				case "String": return "string";
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";
				default: return className(name,params,h); 
			}
		case CAnonymous(fields):
		case CAbstract(name,params):
			switch(name) {
				case "Void": return "";
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";		
				default: return className(name,params,h); 
			}
		}
		return "uintptr"; // if in doubt, its a uintptr
    }
	
    static function haxe2goFuncBody(t:haxe.rtti.CType,h:String):String { // create a dummy go function body
		var gt:String = haxe2goType(t,h);
		switch(gt){
		case "": return " {}";
		case "string": return " string { return \"\"; }";
		case "int","uint","float64": return " " + gt + " { return 0; }" ;
		case "bool": return " bool { return false; }";
		case "interface{}": return " interface{} { return nil; }";
		}
		return " " + gt + " { return 0; }";
    }
	
    static function className(name:String, params:List<haxe.rtti.CType>,h:String):String{ // maybe use a go class name, if declared
					if(isDeclared(name,parser.root))
						return h+d2_(name); 
					else
						return "uintptr";
    }
	
    static function isDeclared(name:String,tree:haxe.rtti.TypeRoot):Bool { // itterate through the decalrations to check that a type is actually declared
    	for(tt in tree)
    		if(isDeclaredHere(name,tt))
    			return true;
    	return false;
    }
	
    static function isDeclaredHere(n:String,tree:haxe.rtti.TypeTree):Bool { // look for the type in this particular spot on the tree
    		switch(tree){
    		case TTypedecl(tt):
    			if(tt.isPrivate) 
  					return false;
				else
					return tt.path==n;
			case TPackage(name,full,subs):
				for(s in subs)
					if(isDeclaredHere(n,s))
						return true;
				return false;
			case TEnumdecl(e):
				if(e.isPrivate)
					return false;
				else
					return e.path==n;
			case TClassdecl(c):
				if(c.isPrivate)
					return false;
				else {
					// TODO check that this code always works correctly
					for(postfix in [""].concat(c.params)){
						var ul = postfix==""?"":".";
						if( (c.path + ul + postfix) ==n )
							return true;
					}
					return false;
				} 
			case TAbstractdecl(a): 
    			if(a.isPrivate)
    				return false;
    			else
	    			return a.path==n;
    		}	
    		return false;
    }
}


class FileUtils // this class modified from http://haxe.org/doc/snip/writetextfilewithbom
{
    /** Write out a text file with correct BOM for UTF-8 encoding */
    public static function writeTextFileWithBOM(text:String,filePath:String):Void
    {
        var f:FileOutput=File.write(filePath,true);
        f.writeByte(239);
        f.writeByte(187);
        f.writeByte(191);
        f.writeString(text);
        f.close();
    }
    
}

