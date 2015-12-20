import React from 'react';
import classNames from 'classnames';

export default class extends React.Component {
    render() {
        return (
            <ul className="menu">
                <li className="button button-cube"><i className="fa fa-rocket"></i>发布</li>
                <li className="button button-cube deploy"><i className="fa fa-chrome"></i>预览</li>
                <li className="button button-circle"><i className="fa fa-floppy-o"></i></li>
                <li className="button button-circle remove"><i className="fa fa-trash"></i></li>
            </ul>
        );
    }
}
