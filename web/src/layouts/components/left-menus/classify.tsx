import React from 'react';
const Classify = (props: any) => {
    return (
        <div>

            {props.data.map((item: any) => {
                return (<div key={item.id}>
                    <div>
                        {item.name}
                    </div>
                    {

                        item.children ? (
                            <div>
                                <Classify data={item.children} />
                            </div>
                        ) : ""
                    }
                </div>)
            })}
        </div>
    )
}
export default Classify