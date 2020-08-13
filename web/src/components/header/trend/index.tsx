import React from 'react';
import './index.less';
import CalendarHeatmap from '@/components/calendar-heatmap';
const Trend = () => {
  return (
    <div className="trend">
      <div className="trend-calendar">
        <div className="trend-calendar-title">动态日历</div>
        <CalendarHeatmap
          startDate={new Date('2020-01-01')}
          endDate={new Date('2020-12-31')}
          values={[
            { date: '2020-1-11', count: 12 },
            { date: '2020-1-12', count: 20 },
            { date: '2020-12-30', count: 1 },
          ]}
        />
      </div>
      <div className="trend-distr">
        {/* 1 */}
        <div className="trend-distr-left">
          <div>发布统计</div>
        </div>
        {/* 2 */}
        <div className="trend-distr-right">
          <div>分类统计图</div>
        </div>
      </div>
    </div>
  );
};
export default Trend;
