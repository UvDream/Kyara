import $axios from '../request';
const ArticleApi = {
  List: (data?: Object) => {
    return $axios.request({
      url: `/article/articleList`,
      data,
      method: 'POST',
    });
  },
  Detail: (params: object) => {
    return $axios.request({
      url: `/article/articleDetail`,
      params,
      method: 'GET',
    });
  },
};
export default ArticleApi;
