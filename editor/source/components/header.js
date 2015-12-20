import React from 'react';
import classNames from 'classnames';

export default class extends React.Component {
    render() {
        const list = this.props.list;
        const actions = this.props.actions;
        return (
            <div id="header">
                <div className="title">构建只为纯粹书写的博客</div>
                <div className="info">
                    <span className="draft">草稿</span>
                    <span className="tag">产品</span>
                    <span className="tag">设计</span>
                </div>
            </div>
        );
    }
}
