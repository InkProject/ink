import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction, editorAction, menuAction } from '../actions';

class Menu extends Component {
    onInboxClick() {
        if (this.props.show) {
            this.props.menuAction.showModal(true)
        } else {
            this.props.listAction.showList()
        }
    }
    render() {
        const show = this.props.show;
        const loading = this.props.loading;
        const listAction = this.props.listAction;
        return (
            <ul id="menu" className="list">
                <li>
                    <button className="button button-circle inbox" onClick={() => this.onInboxClick()}>
                        <i className={classNames('fa', {['fa-'+(show?'plus':'inbox')]: true})}></i>
                    </button>
                </li>
                <li><button className="button button-circle setting"><i className="fa fa-wrench"></i></button></li>
                <li><button className="button button-circle theme"><i className="fa fa-moon-o"></i></button></li>
                <li><button className="button button-circle help"><i className="fa fa-hashtag"></i></button></li>
                <li><i id="loading" className={classNames('fa fa-cog fa-spin', {hide: !loading})}></i></li>
            </ul>
        );
    }
}

export default connect(null, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch),
        menuAction: bindActionCreators(menuAction, dispatch)
    };
})(Menu);
