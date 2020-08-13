import { Effect, Reducer, Subscription } from 'umi';

export interface IndexModelState {
  name: string;
  count: number;
}

export interface IndexModelType {
  namespace: 'index';
  state: IndexModelState;
  effects: {
    query: Effect;
  };
  reducers: {
    save: Reducer<IndexModelState>;
  };
  subscriptions: { setup: Subscription };
}

const IndexModel: IndexModelType = {
  namespace: 'index',
  state: {
    name: '呵呵',
    count: 1,
  },
  reducers: {
    save(state, action) {
      const name = '新的名字';
      return {
        ...state,
        name: name,
        ...action.payload,
      };
    },
    add(state, action) {
      const newCount = state.count + 1;
      return {
        ...state,
        count: newCount,
        ...action.payload,
      };
    },
  },
};

export default IndexModel;
