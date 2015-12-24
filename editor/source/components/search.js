import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction } from '../actions';

export default class Search extends Component {
    render() {
        return (
            <label id="search-wrap" htmlFor="search">
                <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="搜索..." onFocus={this.props.listAction.showList} />
            </label>
        );
    }
}

export default connect(function(state) {
    return state;
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch)
    };
})(Search);
