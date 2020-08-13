import React from 'react';
import { Modal, Button, Tabs } from 'antd';
import "./index.less"
const { TabPane } = Tabs;
const Praise = (props: any) => {
    const handleOk = () => {
        props.changeVisible(false)
    }
    return (
        <div className="praise">
            <Modal
                title="赞赏作者"
                width={300}
                visible={props.visible}
                onOk={handleOk}
                onCancel={handleOk}
                footer={[
                    <div style={{ textAlign: "center" }}>
                        <Button type="primary" shape="round" onClick={handleOk} >
                            确定
                        </Button>
                    </div>
                ]}
            >
                <div className="praise-content">
                    <Tabs defaultActiveKey="1" centered>
                        <TabPane tab="支付宝支付" key="1">
                            <img src="https://gitee.com/Uvdream/images/raw/master/images/20200812100455.png" alt="" />
                        </TabPane>
                        <TabPane tab="微信支付" key="2">
                            <img src="https://gitee.com/Uvdream/images/raw/master/images/20200812100455.png" alt="" />
                        </TabPane>
                    </Tabs>
                </div>
            </Modal>
        </div>
    )
}
export default Praise;