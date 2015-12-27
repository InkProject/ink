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
import Search from './components/search';
import Toolbar from './components/toolbar';

class App extends Component {
    render() {
        const list = this.props.list;
        return (
            <div id="container">
                <Left list={list} listComponent={this.props.List} />
                <Search />
                <Toolbar />
                {this.props.Editor}
            </div>
        );
    }
}

export default connect(function(state) {
    return state;
})(App);
