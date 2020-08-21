import { dynamic } from 'umi';

export default dynamic({
  loader: async function() {
    const { default: HugeA } = await import(/* webpackChunkName: "external_A" */ './markdown.tsx');
    return HugeA;
  },
});
