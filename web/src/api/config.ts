import { baseUrl } from './env';

export default {
  method: 'POST',
  // 基础url前缀
  baseURL: baseUrl,
  // 请求头信息
  headers: {
    'Content-Type': 'application/json;charset=UTF-8',
  },
  // 参数
  data: {},
  // 设置超时时间
  timeout: 30000,
  // 携带凭证
  withCredentials: false,
  // 返回数据类型
  responseType: 'json',
};
