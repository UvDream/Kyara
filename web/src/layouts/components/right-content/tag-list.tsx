import React, { useState } from 'react';
import "./index.less"


const TagList = () => {
    const [tag] = useState([
        { id: 1, name: "前端" },
        { id: 2, name: "web" },
        { id: 3, name: "后端" },
      { id: 4, name: "前端" },
      { id: 5, name: "web" },
      { id: 6, name: "后端" }
    ])
    return (
        <div className="message">
            <div className="message-title">标签云</div>
            <div className="message-tags">
                {tag.map(item => {
                    return (
                        <span key={item.id} className="message-tag">{item.name}</span>
                    )
                })}
            </div>
        </div>
    )
}
export default TagList
