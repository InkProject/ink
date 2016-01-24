import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction, menuAction, modalAction } from '../actions';

class Modal extends Component {
    constructor(props) {
        super(props);
    }
    closeModal() {
        this.props.menuAction.showModal(false);
    }
    createArticle() {
        const name = this.input.value;
        this.props.modalAction.createArticle(name);
        this.props.menuAction.showModal(false);
    }
    render() {
        const menu = this.props.menu;
        return (
            <div id="modal" className={classNames({hide: !menu.get('modal').get('show')})}>
                <div className="overlay" onClick={() => this.closeModal()}></div>
                <div className="content">
                    <div className="title"><i className="fa fa-inbox"></i>创建新文章</div>
                    <input type="text" className="name" placeholder="文件名，建议使用-与字母" ref={(input) => {
                        if (input) {
                            input.focus();
                            this.input = input;
                        }
                    }} />
                <div className="confirm hover" onClick={() => this.createArticle()}><i className="fa fa-inbox"></i>确认</div>
                </div>
            </div>
        );
    }
}

export default connect(function(state) {
    return {
        menu: state.menu
    };
}, function(dispatch) {
    return {
        menuAction: bindActionCreators(menuAction, dispatch),
        modalAction : bindActionCreators(modalAction, dispatch)
    };
})(Modal);
