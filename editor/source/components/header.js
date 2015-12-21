import React from 'react';
import Component from './index';
import classNames from 'classnames';
import _ from 'lodash';

export default class extends Component {
    render() {
        let title = this.props.title || '键入文章标题';
        let tags = [];
        if (_.isArray(this.props.tags))
            tags = this.props.tags;
        return (
            <div id="header">
                <div className="title">{title}</div>
                <div className="info">{
                    tags.map(item =>
                        <span className="tag">{item}</span>
                    )
                }</div>
            </div>
        );
    }
}
