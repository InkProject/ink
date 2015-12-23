import React from 'react';
import Component from './index';
import classNames from 'classnames';

import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { menuAction, editorAction } from '../actions';

class Menu extends Component {
    render() {
        const menu = this.props.menu;
        const menuAction = this.props.menuAction;
        return (
            <ul id="menu" className="menu">
                <li className="button button-circle inbox" onClick={menuAction.showList}>
                    <i className={classNames('fa', {['fa-'+(menu.get('show')?'plus':'inbox')]: true})}></i>
                </li>
                <li className="button button-circle setting"><i className="fa fa-wrench"></i></li>
                <li className="button button-circle theme"><i className="fa fa-moon-o"></i></li>
                <li className="button button-circle help"><i className="fa fa-hashtag"></i></li>
                <li><i id="loading" className={classNames('fa fa-cog fa-spin', {hide: !menu.get('loading')})}></i></li>
            </ul>
        );
    }
}

export default connect(null, function(dispatch) {
    return {
        menuAction: bindActionCreators(menuAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch)
    };
})(Menu);
