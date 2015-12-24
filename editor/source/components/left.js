import React from 'react';
import Component from './index';
import classNames from 'classnames';
import List from './list';
import Menu from './menu';

export default class Left extends Component {
    render() {
        const show = this.props.show;
        const loading = this.props.loading;
        return (
            <div id="left" className={classNames({close: !show})}>
                <Menu show={show} loading={loading} />
                {this.props.listComponent}
            </div>
        );
    }
}
