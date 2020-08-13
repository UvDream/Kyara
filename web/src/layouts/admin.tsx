import React from 'react';

export default (props: any) => {
    return (
        <>

            管理页面
            <div style={{ padding: 20 }}>{props.children}</div>
        </>
    );
}