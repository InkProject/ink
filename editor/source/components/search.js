import React from 'react';
import classNames from 'classnames';

export default class extends React.Component {
    render() {
        return (
            <label id="search-wrap" htmlFor="search">
                <i className="fa fa-search"></i>
                <input id="search" type="text" placeholder="搜索..." />
            </label>
        );
    }
}
