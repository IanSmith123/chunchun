set ws=WScript.CreateObject("WScript.Shell")
path = createobject("Scripting.FileSystemObject").GetFile(Wscript.ScriptFullName).ParentFolder.Path
ws.Run path+"/chunchun.exe",0