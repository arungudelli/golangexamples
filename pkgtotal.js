function getAllPackages(){
    var pkgdir = document.getElementsByClassName("pkg-dir")[0];
    var trs = pkgdir.getElementsByTagName("tr");
    var totalpackages = [];
    var currentPackage = "";

    for(var i=0;i<trs.length;i++){
        if(i==0) continue;
        var tr = trs[i];
        var an = tr.getElementsByTagName("a")[0];
        
        var Name = an.innerText;
        var Desc = tr.getElementsByClassName("pkg-synopsis")[0].innerText;
        var Link = an.href;
        totalpackages.push({
            Name:Name,
            Desc:Desc,
            Link:Link
        });
        
    }

return totalpackages;
}

function getpkgInfo(){
    var desc = [];
    var overview=document.getElementById("pkg-overview");
    if(!overview) return JSON.stringify(desc);
    var overviewDesc = overview.getElementsByClassName("expanded")[0];
    var childs=overviewDesc.childNodes;
    var desc = [];
    for(var i=0;i<childs.length;i++){
        var child = childs[i];
        if(child.nodeName =="#text") continue;
        desc.push([child.nodeName,child.innerText]);
    }
    return JSON.stringify(desc);
}

function getfunctionDesc(){
    var functionDesc = [];
    var h2s= document.getElementsByClassName("container")[0].childNodes;
    var h2id = [];
    
    for(var i=0;i<h2s.length;i++){
        if(h2s[i].tagName == "H2" && h2s[i].innerText.startsWith("func")){
            h2id.push({
                id:i,
                name:h2s[i].innerText
            });
        }
    }

    if(h2id.length == 0) return JSON.stringify(functionDesc);

    for(var j=0;j<h2id.length;j++){
        
        if(j+1 >= h2id.length) break; 

        var name = h2id[j].name;

        var h2index = h2id[j].id;

        var h2End = h2id[j+1].id;   
        
        var funcd = [];

        for(var i=h2index;i<h2End;i++){
            var hchild = h2s[i];
            if(hchild.nodeName =="#text") continue;
            funcd.push([hchild.nodeName,hchild.innerText]);

        }

        var egId = "example"+"_"+name.replace("func","").replace(" ","");

        var eg = document.getElementById(egId);
        var egText = "";

        if(eg){
            egText = eg.getElementsByClassName("code")[0].value;
        }

        functionDesc.push({
            fname:name,
            fdesc:funcd,
            example:egText
        });

        

    }

    var lastMethod = h2id[h2id.length-1];
    var lname = lastMethod.name;


    var lfuncd = [];

        for(var i=lastMethod.id;i<h2s.length;i++){
            var hchild = h2s[i];
            
            if(h2s[i+1] && h2s[i+1].tagName && h2s[i+1].tagName.startsWith("H")){
                break;
            }
            if(hchild.nodeName =="#text") continue;
            lfuncd.push([hchild.nodeName,hchild.innerText]);

            
        }

        var legId = "example"+"_"+lname.replace("func","").replace(" ","");

        var leg = document.getElementById(legId);
        var legText = "";

        if(leg){
            legText = leg.getElementsByClassName("code")[0].value;
        }

        functionDesc.push({
            Fname:lname,
            Fdesc:lfuncd,
            Example:legText
        });

    return JSON.stringify(functionDesc);
}


function getfunction(){
    var h2s= document.getElementsByClassName("container")[0].childNodes;
    var h2id = [];
    for(var i=0;i<h2s.length;i++){
        if(h2s[i].tagName == "H2" && h2s[i].innerText.startsWith("func")){
            h2id.push({
                id:i,
                name:h2s[i].innerText
            });
        }
    }
    return h2id;
}