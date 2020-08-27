import React, { useEffect } from 'react';
import Vditor from 'vditor'
import "../../../../styles/vditor.css"
import "../../../../styles/vditor.less"
const MarkDown = (props: Object) => {
  const parseDom = (arg: any) => {
    let objE = document.createElement("div");
    objE.innerHTML = arg;
    return objE.childNodes;
  }

  useEffect(() => {
    Vditor.preview(document.getElementById("htmlView"), props.content, {
      speech: {
        enable: true,
      },
      anchor: true,
      hljs: {
        enable: true,
        lineNumber: true,
        style: "native"
      },
      markdown: {
        toc: true
      },
      transform: (html: string) => {
        let arr: { children: { id: (element: ChildNode, id: any, firstChild: ChildNode | null) => any; title: ChildNode | null; children: never[]; }[]; }[] = []
        parseDom(html).forEach(element => {
          if (element.nodeName == "H1" || element.nodeName == "H2" || element.nodeName == "H3") {
            let obj = { id: element.id, title: element.firstChild, children: [] }
            if (element.nodeName == "H1") {
              arr.push(obj)
            }
            if (element.nodeName == "H2") {
              arr[arr.length - 1].children.push(obj)
            }
            if (element.nodeName == "H3") {
              arr[arr.length - 1].children[arr[arr.length - 1].children.length - 1].children.push(obj)
            }
          }
        })
        console.log("目录树结构", arr)
        return html
      }
    })
  }, [props.content])
  return (
    <>
      <div id="htmlView" className="detail-content"></div>
    </>
  )
}
export default MarkDown;
