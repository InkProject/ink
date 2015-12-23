import React from 'react';
import Component from './index';
import classNames from 'classnames';

export default class extends Component {
    render() {
        return (
            <label id="search-wrap" htmlFor="search">
                <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="搜索..." onFocus={this.props.onFocus.bind(this)} />
            </label>
        );
    }
}
