import $axios from '../request';
const ArticleApi = {
  Detail: (data?: Object) => {
    return $axios.request({
      url: `/article/articleList`,
      data,
      method: 'POST',
    });
  },
};
export default ArticleApi;
