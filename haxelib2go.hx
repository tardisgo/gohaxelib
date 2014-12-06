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
        out+=xmlFile("api-xml/cross.xml","cross","X");
        /*
        out+=xmlFile("api-xml/cpp.xml","cpp","P");
        out+=xmlFile("api-xml/cs.xml","cs","R");
        out+=xmlFile("api-xml/flash9.xml","flash9","F");
        out+=xmlFile("api-xml/java.xml","java","J");
        out+=xmlFile("api-xml/js.xml","js","S");
        out+=xmlFile("api-xml/neko.xml","neko","N");
        out+=xmlFile("api-xml/php.xml","php","H");
        */
        FileUtils.writeTextFileWithBOM("package _haxeapi\n//import \"github.com/tardisgo/tardisgo/tardisgolib/hx\"\n"+
        	out,"_haxeapi/def.go");
        
		/* OpenFL */
		//out=xmlFile("api-xml/openfl.xml","","X");
        //FileUtils.writeTextFileWithBOM("package _openfl\n"+out,"_openfl/def.go");
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
    			//if(tt.isPrivate) // may need private types for inherritance
  				//	"";
				//else
				//	"type "+h+d2_(tt.path)+" "+haxe2goType(tt.type,h)+" // TTypedecl\n"; //TODO extend
				"type "+h+d2_(tt.path)+" uintptr /* TTypedecl -  "+tt.type+" */\n"; 

			case TPackage(name,full,subs): 
				var ret:String = "";
				for(s in subs)
					ret += tt(s,h);
				ret;
			case TEnumdecl(e):
				//if(e.isPrivate)
				//	"";
				//else
				//	"type "+h+ d2_(e.path) + " struct{} // TEnumdecl\n"; //TODO extend
				"type "+h+d2_(e.path)+" uintptr /* Enumdecl - "+e.constructors+" */\n"; 
			case TClassdecl(c) :
				if(!c.isPrivate){
					var ret:String = "" ;
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
								//ret += haxe2goFuncBody(_ret,h);
								ret += "\n\n";
							default:
								//ret += "var "+h+d2_(c.path)+"_"+d2_(s.name)+" "+haxe2goType(s.type,h)+"\n";
								// add getters & setters for static vars
								ret += "func "+h+ d2_(c.path) + "_"+d2_(s.name) +"_goget() ";
								//ret += haxe2goFuncBody(s.type,h);
								ret += "\n\n";
								ret += "func "+h+ d2_(c.path) + "_"+d2_(s.name) +"_goset(";
								ret += haxe2goType(s.type,h)+")";
								ret += "\n\n";
							}
						}
					}
					for(postfix in [""].concat(c.params)){
						var ul = postfix==""?"":"_";
						var typName:String = h+ d2_(c.path) + ul + postfix ;
						var scName:String = "uintptr /* should be: hx.Dynamic */" ;
						if(c.superClass != null ) {
							scName= h+d2_(c.superClass.path) ;
							trace(scName);
						}

						ret += "type "+ typName+ " "+scName+" // TClassdecl\n\n";
						//ret += "type "+ typName+ " struct { // TClassdecl\n";
						//for(m in c.fields)
						//	if(m.isPublic)		
						//		switch(m.type){
						//		case CFunction(args,_ret): 
						//			//NoOp
						//		default:
						//			var vt:String = haxe2goType(m.type,h);
						//			switch(vt){
						//			case "bool","string","int","uint","float64","uintptr":
						//				// NoOp
						//			default:
						//				vt = "/*pseudo*/ *"+vt;
						//			} 
						//			ret += "\t"+h+d2_(m.name) + " " + vt +"\n";
						//		}				
						//ret += "}\n";
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
										ret += "" + h+a.name + " " + haxe2goType(a.t,h);
										hadFirst=true;
									}	
								}
								if(args.length>0) 
									ret += "";
								ret += ")";
								if(m.name=="new")
									ret+=" "+h+d2_(c.path) ; //+" { return "+h+d2_(c.path)+"{}; } ";
								//else
								//	ret+=haxe2goFuncBody(_ret,h);
								ret += "\n\n";
							default:
								ret += "func (x "+h+ d2_(c.path) + ") "+h+d2_(m.name) +"_goget() ";
								//ret += haxe2goFuncBody(m.type,h);
								ret += "\n\n";
								ret += "func (x "+h+ d2_(c.path) + ") "+h+d2_(m.name) +"_goset(";
								ret += haxe2goType(m.type,h)+")";
								// ret += "{}"
								ret += "\n\n";
							}
						}
					ret;
				} else "";
			
			case TAbstractdecl(a): 
    			//if(a.isPrivate)
    			//	return ""
    			//else {
	    		//	return "type "+h+ d2_(a.path) + " struct{} // TAbstractdecl\n";
    			//}
    			"type "+h+d2_(a.path)+" uintptr /* TAbstractdecl - "+a.athis+" */\n";
    		}
    }
	
    static function d2_(x:String):String { // convert dots to underlines, underline to tripple underline
    	if(x==null)
    		return null;
    	if(x.length>2) {
	    	var r:String = x.split("_").join("...").split(".").join("_"); // transform orig "_" to "___" 
	    	r = r.split("<").join("XltX").split(">").join("XgtX"); // TODO consider how to deal with these chars
	    	return r;
    	} else
    		return x;
    }

	
    static function haxe2goType(t:haxe.rtti.CType,h:String):String { // mapping of haxe to Go types
    	// TODO work out which of "Void" or Int/Float/Bool is actually the one used...
		switch(t){
		case CUnknown:
			return "uintptr"; // if in doubt, its a uintptr
		case CTypedef(name,params):
			switch(name) {
				case "Void": return "/*Void*/"; 
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";
				case "Null": return "uintptr";
				default: return className(name,params,h); 
			}
		case CFunction(args,_ret):
			return "func ()"; // TODO should be: "func (args...) ret_"
		case CEnum(name,params):
			switch(name) {
				case "Void": return "/*Void*/";
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";
				case "Null": return "uintptr";
				default: return className(name,params,h); 
			}
		case CDynamic(t):
			return "uintptr"; 
		case CClass(name,params):
			switch(name){
				case "Void": return "/*Void*/";
				case "String": return "string";
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";
				case "Null": return "uintptr";
				default: return className(name,params,h); 
			}
		case CAnonymous(fields):
			var r:String = "struct { "; // TODO should list internal fields here:
			for(f in fields)
				r = r + "_" + f.name + " " + haxe2goType(f.type,h) + "; ";
			return r+"}";
		case CAbstract(name,params):
			switch(name) {
				case "Void": return "/*Void*/";
				case "Bool": return "bool";
				case "Int": return "int";
				case "UInt": return "uint";
				case "Float": return "float64";		
				case "Null": return "uintptr";
				default: return className(name,params,h); 
			}
		}
		return "uintptr"; // if in doubt, its a uintptr
    }
	
    static function haxe2goFuncBody(t:haxe.rtti.CType,h:String):String { // create a dummy go function body
		var gt:String = haxe2goType(t,h);
		switch(gt){
		case "/*Void*/": return "/*{}*/\n\n";
		case "string": return " string /*{ return \"\"; }*/\n\n";
		case "int","uint","float64": return " " + gt + " /*{ return 0; }*/\n\n" ;
		case "bool": return " bool /*{ return false; }*/\n\n";
		case "interface{}": return " interface{} /*{ return nil; }*/\n\n";
		case "uintptr": return " uintptr /*{ return 0; }*/\n\n";
		}
		return " " + gt + " /*{ return "+gt+"{}; }*/\n\n";
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

