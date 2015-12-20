import 'font-awesome/less/font-awesome.less';
import './styles/index.less';

import React from 'react';
import ReactDom from 'react-dom';
import classNames from 'classnames';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { listAction } from './actions';

import Editor from './components/editor';
import List from './components/list';
import Menu from './components/menu';
import Header from './components/header';
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends React.Component {
    componentWillMount() {
        this.props.listAction.fetchList();
    }
    render() {
        const menu = this.props.menu;
        const list = menu.get('list');
        const editor = this.props.editor;
        const listAction = this.props.listAction;
        return (
            <div id="container">
                <div id="left" onMouseLeave={listAction.hideList} onMouseOver={listAction.showList}>
                    <Menu menu={menu} />
                    <div id="files" className={classNames({hide: !menu.get('show')})}>
                        <Search />
                        <List list={list} onOpenArticle={listAction.openArticle} />
                    </div>
                </div>
                <div id="right">
                    <Toolbar />
                </div>
                <Header />
                <Editor content={editor.get('content')} />
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch)
    };
})(App);
