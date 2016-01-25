import React from 'react';
import Component from './index';
import classNames from 'classnames';
import List from './list';
import Menu from './menu';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction } from '../actions';

export default class Left extends Component {
    render() {
        const list = this.props.list;
        const show = list.get('show');
        const loading = list.get('loading');
        return (
            <div id="left" className={classNames({close: !show})}>
                <Menu show={show} loading={loading} />
                <List />
            </div>
        );
    }
}

export default connect(function(state) {
    return {
        list: state.list
    }
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(Left);
