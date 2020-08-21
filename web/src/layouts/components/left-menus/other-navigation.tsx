import React, { useState } from 'react';
import "./index.less"
// @ts-ignore
import Classify from "./classify.tsx"
const OtherNavigation = () => {
    const [menus] = useState([{
        id: 1, name: "分类", icon: "",open:true, children: [
            { id: "1-1", name: "前端开发", count: "100",open:true ,children: [
                { id: "1-1", name: "html", count: "100" },
                { id: "1-2", name: "javascript", count: "100" }
              ]},
            { id: "1-2", name: "后端开发", count: "100" }
        ]
    }])
    return (
        <div className="nav">
            <div className="nav-title">其他</div>
            <div >
                <Classify data={menus} />
            </div>
        </div>
    )
}
export default OtherNavigation
