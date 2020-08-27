import React, { useState, useEffect } from 'react';
import { IndexModelState, ConnectRC, connect } from 'umi';
import './index.less';
import Praise from "./components/praise"
import { Button } from 'antd';
import { createFromIconfontCN } from '@ant-design/icons';
// import MarkDown from "./components/markdown"
import AsyncMarkdown from "./components/async-markdown"
import ArticleApi from "@/api/blog/index"
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
interface PageProps {
    index: IndexModelState;
    loading: boolean;
}
const Detail = (props: any) => {
    const [markdownContent, setMarkdown] = useState()
    const [htmlContent, setHtml] = useState()
    const [visible, setVisible] = useState(false)
    useEffect(() => {
        if (props?.data?.msg) {
            setMarkdown(props.data.msg.article_content)
            setHtml(props.data.msg.article_html)
        }
    }, [markdownContent, htmlContent])

    return <div className="detail" >
        <div className="detail-title">
            <div className="detail-title-top">
                <h1>这是文章的标题</h1>
                <div>
                    <IconFont type="uv-caret-forward-circle-outline" className="detail-title-top-icon" />
                    {/* <IconFont type="uv-pause-circle-outline" /> */}
                    <IconFont type="uv-book-outline" className="detail-title-top-icon" />
                    <IconFont type="uv-print-outline" className="detail-title-top-icon" />
                </div>
            </div>
            <div>
                <IconFont
                    type="uv-person-outline"
                    className="list-normal-bottom-iconO"
                />
                <span>作者</span>
                <IconFont
                    type="uv-time-outline"
                    className="list-normal-bottom-icon"
                />
                <span>2020年10月1日</span>
                <IconFont
                    type="uv-chatbubble-outline"
                    className="list-normal-bottom-icon"
                />
                <span>1220条评论</span>
                <IconFont
                    type="uv-eye-outline"
                    className="list-normal-bottom-icon"
                />
                <span>1220次浏览</span>
                <IconFont
                    type="uv-information-outline"
                    className="list-normal-bottom-icon"
                />
                <span>122字数</span>
            </div>
        </div>
        <div style={{ backgroundImage: `url(https://gitee.com/Uvdream/images/raw/master/images/20200806150710.png)` }} className="detail-img"></div>
        <AsyncMarkdown content={markdownContent} />
        <div dangerouslySetInnerHTML={{ __html: htmlContent }}></div>
        <div className="detail-praise">
            <div className="detail-praise-time">
                <section>
                    <IconFont
                        type="uv-time-outline"

                    />
                    最后修改时间:2020年10月1日</section>
                <section>
                    <IconFont
                        type="uv-alert-circle-outline"

                    />
                    允许转载</section>
            </div>
            <div className="detail-praise-text">
                <Button type="primary" danger shape="round" onClick={() => { setVisible(true) }} >
                    <IconFont type="uv-wallet-outline" />
                    赞赏</Button>
                <p>如果觉得我的文章对你有用，请随意赞赏</p>
            </div>
        </div>
        <div className="detail-btn">
            <Button shape="round" >
                上一篇</Button>
            <Button shape="round" >
                下一篇</Button>
        </div>
        <Praise visible={visible} changeVisible={() => { setVisible(false) }} />
        {/* <div>
            Hello {name}
            <p onClick={() => { dispatch({ type: 'index/add' }) }}>{count}</p>
        </div> */}
    </div>;
}
Detail.getInitialProps = (async (ctx: any) => {
    let msg = ""
    console.log("获取", ctx)
    let res = await ArticleApi.Detail({ id: ctx.history.location.query.id })
    if (res.code == 0) {
        msg = res.data
        return Promise.resolve({
            data: {
                title: '文章详情页',
                msg: msg
            }
        })
    }

})
export default Detail