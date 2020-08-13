import React from 'react';
import "./index.less"
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const BlogMessage = () => {
    return (
        <div className="message">
            <div className="message-title">
                博客信息
            </div>
            <div className="message-content">
                <section>
                    <div>
                        <IconFont type="uv-ribbon-outline" className="message-content-icon" />
                        文章数目
                    </div>
                    <div className="message-tag">
                        200
                    </div>
                </section>
                <section>
                    <div>
                        <IconFont type="uv-chatbubble-outline" className="message-content-icon" />
                        评论数量
                    </div>
                    <div className="message-tag" >
                        100
                    </div>
                </section>
                <section>
                    <div>
                        <IconFont type="uv-browsers-outline" className="message-content-icon" />
                        运行天数
                    </div>
                    <div className="message-tag">
                        4年7天
                    </div>
                </section>
                <section>
                    <div>
                        <IconFont type="uv-pulse-outline" className="message-content-icon" />
                        最后活动
                    </div>
                    <div className="message-tag">4天前</div>
                </section>
            </div>
        </div>
    )
}
export default BlogMessage