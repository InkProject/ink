import React from 'react';
import Component from './index';
import classNames from 'classnames';
import _ from 'lodash';

export default class extends Component {
    render() {
        let title = this.props.title || '键入文章标题';
        var tags = [];
        if (_.isArray(this.props.tags) && this.props.tags.length > 0)
            tags = this.props.tags;
        return (
            <div id="header" className={classNames({edit: this.props.edit})}>
                <div className="title hover" onClick={this.props.onClick}><i className="fa fa-cog"></i><span>{title}</span></div>
                <div className="info">{
                    tags.map(item =>
                        item ? <span className="tag">{item}</span> : null
                    )
                }</div>
            </div>
        );
    }
}
