import React, { useState } from 'react';
import "./index.less"
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const Navigation = () => {
    const [menus] = useState([
        { title: "首页", id: 1, icon: "uv-home-outline" },
        { title: "仓库", id: 2, icon: "uv-git-merge-outline" },
        { title: "相册", id: 6, icon: "uv-image-outline" },
        { title: "归档", id: 3, icon: "uv-file-tray-full-outline" },
        { title: "留言", id: 4, icon: "uv-at-sharp" },
        { title: "关于", id: 5, icon: "uv-cafe-outline" },

    ])
    return (
        <nav className="nav">
            <div className="nav-title">导航</div>
            <ul>
                {
                    menus.map(item => {
                        return (
                            <li key={item.id}>
                                <IconFont type={item.icon} className="nav-icon" />
                                <span>{item.title}</span>
                            </li>
                        )
                    })
                }
            </ul>
        </nav>
    )
}
export default Navigation;