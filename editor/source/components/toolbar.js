import React from 'react';
import Component from './index';
import classNames from 'classnames';

export default class extends Component {
    constructor(props) {
        super(props);
        this.state = {confirm: false};
    }
    showConfirm(flag) {
        this.setState({confirm: flag});
    }
    onRemoveClick() {
        this.setState({confirm: false});
        alert(1);
    }
    render() {
        return (
            <ul id="right">
                <li><button className="button button-cube"><i className="fa fa-rocket"></i>发布</button></li>
                <li><button className="button button-cube deploy"><i className="fa fa-chrome"></i>预览</button></li>
                <li><button className="button button-circle"><i className="fa fa-floppy-o"></i></button></li>
                <li><button className="button button-circle remove" onFocus={() => this.showConfirm(true)} onBlur={() => this.showConfirm(false)}><i className="fa fa-trash"></i></button></li>
                {this.state.confirm ? <div id="confirm" className="hover" onMouseDown={() => this.onRemoveClick()}>确认删除</div> : null}
            </ul>
        );
    }
}
