import React, { useState, useEffect } from 'react';
import '@/components/calendar-heatmap/styles.css';
import { initTheme } from '@/utils/theme';
import './index.less';
import List from './components/list';
import { Pagination } from 'antd';
import ArticleApi from "@/api/blog/index"
const Home = (props: any) => {
  const [articleList, setArticleList] = useState([]);
  useEffect(() => {
    initTheme(true);
    getData();
  }, []);
  const getData = async () => {
    let res = await ArticleApi.Detail({
      "search": "",
      "page": 1,
      "pageSize": 5
    })
    if (res.data.code == 0) {
      let msg = res.data.data.msg
      setArticleList(msg)
    }

  }

  const themeChange = (data: Boolean) => {
    initTheme(data);
  };
  return (
    <div className="home">
      {
        articleList.length > 0 ? (articleList.map((item: any) => {
          console.log(item)
          return (
            <List key={item.ID} data={item} />
          )
        })) : "暂无数据"
      }
      {/* <List data={data} />
      <List data={data1} />
      <List data={data2} />
      <List data={data3} />
      <List data={data4} /> */}
      <div className="home-page">
        <section>
          <Pagination size="small" defaultCurrent={1} total={50} />
        </section>
      </div>
      {/* <button
        onClick={() => {
          themeChange(true);
        }}
      >
        切换
      </button>
      <button
        onClick={() => {
          themeChange(false);
        }}
      >
        切换
      </button>
      博客首页 */}
    </div>
  );
};
export default Home;
