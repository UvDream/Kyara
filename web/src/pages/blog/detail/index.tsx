import React, { useState, useEffect } from 'react';
import { IndexModelState, ConnectRC, connect } from 'umi';
import './index.less';
import Vditor from 'vditor'
import "../../../styles/vditor.css"
import "../../../styles/vditor.less"
import Praise from "./components/praise.tsx"
import { Button } from 'antd';
import { createFromIconfontCN } from '@ant-design/icons';
const IconFont = createFromIconfontCN({
    scriptUrl: '//at.alicdn.com/t/font_1762601_mkiqfjs1yth.js',
});
interface PageProps {
    index: IndexModelState;
    loading: boolean;
}

const Detail: ConnectRC<PageProps> = ({ index, dispatch }) => {
    useEffect(() => {
        Vditor.preview(document.getElementById("htmlView"), markdownContent, {
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
                    }
                })
                console.log(arr)
                return html
            }
        })
    }, [])
    const { name } = index;
    const { count } = index;
    const [visible, setVisible] = useState(false)
    const parseDom = (arg: any) => {
        let objE = document.createElement("div");
        objE.innerHTML = arg;
        return objE.childNodes;
    }
    const editTitle = () => {
        dispatch({
            type: "index/save"
        })
    }
    const [markdownContent] = useState('# P01:课程介绍和环境搭建\n' +
        '[ **M** ] arkdown + E [ **ditor** ] = **Mditor**  \n' +
        '> Mditor 是一个简洁、易于集成、方便扩展、期望舒服的编写 markdown 的编辑器，仅此而已... \n\n' +

        '**这是加粗的文字**\n\n' +
        '*这是倾斜的文字*`\n\n' +
        '***这是斜体加粗的文字***\n\n' +
        '~~这是加删除线的文字~~ \n\n' +
        '\`console.log(111)\` \n\n' +
        '# p02:来个Hello World 初始Vue3.0\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n' +
        '***\n\n\n' +
        '# p03:Vue3.0基础知识讲解\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n\n' +
        '# p04:Vue3.0基础知识讲解\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n\n' +
        '#5 p05:Vue3.0基础知识讲解\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n\n' +
        '# p06:Vue3.0基础知识讲解\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n\n' +
        '# p07:Vue3.0基础知识讲解\n' +
        '> aaaaaaaaa\n' +
        '>> bbbbbbbbb\n' +
        '>>> cccccccccc\n\n' +
        '``` var a=11; ```'
    )
    return <div className="detail" onClick={editTitle}>
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
        <div id="htmlView" className="detail-content"></div>
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
};
export default connect(({ index }: { index: IndexModelState }) => ({
    index,
}))(Detail);