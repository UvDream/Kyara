import React,{useState} from 'react';
import "./index.less"
import { RightCircleOutlined ,DownCircleOutlined} from "@ant-design/icons"

const Classify = (props: any) => {
  const [data,setData]=useState(props.data)
  const menuClick=(item:Object)=>{
    item.open=!item.open
    setData(data.slice())
  }
    return (
        <>
            {data.map((item: any) => {
                return (<div key={item.id} className="menu">
                    <div className="menu-list"   onClick={()=>{menuClick(item)}}>
                      <div className="menu-list-icon">
                        {
                          item.children&&item.open?<DownCircleOutlined   />:""
                        }
                        {
                          item.children&&!item.open?<RightCircleOutlined  />:""
                        }
                      </div>
                       <div>
                         <span className="menu-list-text">{item.name}</span>
                        {item.count? <span className="message-tag">{item.count}</span>:""}
                       </div>
                    </div>
                  {

                    item.children &&item.open? (
                      <div className="menu-children">
                        <Classify data={item.children}   />
                      </div>
                    ) : ""
                  }

                </div>)
            })}
        </>
    )
}
export default Classify
