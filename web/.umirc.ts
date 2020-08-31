import { defineConfig } from 'umi';
export default defineConfig({
  ssr: {
    // forceInitial: true,
  },
  dva: {
    immer: true,
  },
  antd: {},
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    {
      path: '/index',
      component: '@/layouts/index',
      routes: [
        {
          path: '/index/home',
          title: '首页',
          component: '@/pages/blog/home/index.tsx',
        },
        {
          path: '/index/detail',
          title: '文章详情页',
          component: '@/pages/blog/detail/index.tsx',
        },
      ],
    },
    {
      path: '/admin',
      title: '后台管理',
      component: '@/layouts/admin',
      routes: [{ path: '/admin/index', component: '@/pages/admin/index.tsx' }],
    },
  ],
});
