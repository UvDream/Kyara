import React, { useState, useEffect } from 'react';
import '@/components/calendar-heatmap/styles.css';
import { initTheme } from '@/utils/theme';
import { IGetInitialProps } from 'umi';
import './index.less';
import List from './components/list';
import { Pagination } from 'antd';
import ArticleApi from "@/api/blog/index"
const Home = (props: any) => {
  const [articleList, setArticleList] = useState([]);
  useEffect(() => {
    initTheme(true);
    if (props?.data?.articleList) {
      setArticleList(props.data.articleList)
    }
  }, []);


  const pageChange = (page: number, pageSize: number) => {
  }
  return (
    <div className="home">
      {
        articleList.length > 0 ? (articleList.map((item: any) => {
          return (
            <List key={item.ID} data={item} />
          )
        })) : "暂无数据"
      }

      <div className="home-page">
        <section>
          <Pagination size="small" defaultCurrent={1} total={50} />
        </section>
      </div>

    </div>
  );
};

Home.getInitialProps = (async (ctx) => {
  let msg
  let res = await ArticleApi.List({
    "search": "",
    "page": 1,
    "pageSize": 5
  })
  if (res.code == 0) {
    msg = res.data.msg
    return Promise.resolve({
      data: {
        title: 'Hello World',
        articleList: msg
      }
    })
  }


}) as IGetInitialProps;
export default Home;
