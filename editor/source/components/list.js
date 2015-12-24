import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { Link } from 'react-router';
import { listAction, editorAction } from '../actions';

class List extends Component {
    componentWillMount() {
        this.props.listAction.fetchList();
    }
    render() {
        const list = this.props.list;
        const id = this.props.params.id;
        if (id) {
            this.props.listAction.openArticle(id);
        }
        return (
            <ul id="list" className={classNames({hide: this.props.hide})}>{
                list.get('data').map(item =>
                    <li className={classNames('item hover', {selected: item.get('id') == id})} key={item.get('id')} >
                        <Link to={`/edit/${item.get('id')}`}>
                            <div className="head">
                                <span className="date">2015-03-01 18:00</span>
                            </div>
                            <div className="title">{item.get('title')}</div>
                            <div className="preview">{item.get('preview')}</div>
                        </Link>
                    </li>
                )
            }</ul>
        );
    }
}

export default connect(function(state) {
    return {
        list: state.list
    };
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(List);
