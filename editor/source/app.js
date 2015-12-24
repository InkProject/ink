import 'font-awesome/less/font-awesome.less';
import './styles/index.less';

import React from 'react';
import Component from './components';
import ReactDom from 'react-dom';
import classNames from 'classnames';
import { bindActionCreators } from 'redux';
import { connect } from 'react-redux';
import { editorAction } from './actions';

import Left from './components/left';
import Editor from './components/editor';
import Header from './components/header';
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends Component {
    render() {
        const list = this.props.list;
        const editor = this.props.editor;
        return (
            <div id="container">
                <Left show={list.get('show')} loading={list.get('loading')} listComponent={this.props.children} />
                <Search />
                <Toolbar />
                <Header title={editor.get('title')} tags={editor.get('tags')}/>
                <Editor content={editor.get('content')} />
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
})(App);
