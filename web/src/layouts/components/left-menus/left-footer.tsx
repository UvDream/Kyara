import React from 'react'
import "./index.less"
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
const LeftFooter = () => {
    return (
        <footer className="footer">
            <div>
                <div>
                    <IconFont type="uv-stats-chart-outline" />
                </div>
                <section>管理</section>
            </div>
            <div>
                <div> <IconFont type="uv-newspaper-outline" /></div>
                <section>文章</section>
            </div>
            <div>
                <div> <IconFont type="uv-chatbubble-ellipses-outline" /></div>
                <section>评论</section>
            </div>
        </footer>
    )
}
export default LeftFooter