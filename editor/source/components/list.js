import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { menuAction, editorAction } from '../actions';

class List extends Component {
    render() {
        const list = this.props.list;
        return (
            <ul id="list" className={classNames({hide: this.props.hide})}>{
                list.get('data').map(item =>
                    <li className={
                        classNames('item hover', {
                            selected: item.get('id') == list.get('selected').get('id')})
                        } key={item.get('id')} onClick={this.props.menuAction.openArticle.bind(this, item.get('id'))}>
                        <div className="head">
                            <span className="date">2015-03-01 18:00</span>
                        </div>
                        <div className="title">{item.get('title')}</div>
                        <div className="preview">{item.get('preview')}</div>
                    </li>
                )
            }</ul>
        );
    }
}

export default connect(null, function(dispatch) {
    return {
        menuAction: bindActionCreators(menuAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(List);
