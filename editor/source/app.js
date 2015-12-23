import 'font-awesome/less/font-awesome.less';
import './styles/index.less';

import React from 'react';
import Component from './components';
import ReactDom from 'react-dom';
import classNames from 'classnames';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { menuAction, editorAction } from './actions';

import Editor from './components/editor';
import List from './components/list';
import Menu from './components/menu';
import Header from './components/header';
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends Component {
    componentWillMount() {
        this.props.menuAction.fetchList();
    }
    render() {
        const menu = this.props.menu;
        const editor = this.props.editor;
        const menuAction = this.props.menuAction;
        const list = menu.get('list');
        return (
            <div id="container">
                <div id="left" className={classNames({close: !menu.get('show')})}>
                    <Menu menu={menu} />
                    <List list={list} hide={!menu.get('show')} />
                </div>
                <Search onFocus={menuAction.showList} />
                <Toolbar />
                <Header title={editor.get('title')} tags={editor.get('tags')}/>
                <Editor content={editor.get('content')} />
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
}, function(dispatch) {
    return {
        menuAction: bindActionCreators(menuAction, dispatch)
    };
})(App);
