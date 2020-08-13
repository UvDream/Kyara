import React, { useState } from 'react';
import "./index.less"
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const User = () => {
    const [songListShow, setSongListShow] = useState(false)
    return (
        <div className="user">
            <IconFont type="uv-reorder-four-outline" style={{ fontSize: "18px" }} />
            <IconFont type="uv-chatbox-ellipses-outline" style={{ fontSize: "18px" }} onClick={() => { setSongListShow(!songListShow) }} />
            <IconFont type="uv-person-outline" style={{ fontSize: "18px" }} />
            {
                songListShow ? (
                    <div className="user-song swing-bottom-bck">
                        歌词列表
                    </div>
                ) : ""
            }
        </div>
    )
}
export default User;