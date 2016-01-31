import React from 'react'
import Component from './index'
import classNames from 'classnames'
import ace from 'brace'
import 'brace/mode/markdown'
import 'brace/theme/tomorrow'
import _ from 'lodash'

import { bindActionCreators } from 'redux'
import { connect } from 'react-redux'
import { listAction, editorAction, utilAction } from '../actions'

import Header from './header'

class Editor extends Component {
    constructor(props) {
        super(props)
        this.state = {
            configMode: false,
            toolbar: {
                show: false,
                selectMode: false
            }
        }
    }
    switchConfigMode() {
        this.setState({configMode: !this.state.configMode})
        if (this.state.configMode) {
            this.contentEditor.focus()
            this.props.utilAction.showTip('auto', '切换至内容')
        } else {
            this.configEditor.focus()
            this.props.utilAction.showTip('auto', '切换至配置')
        }
        this.setEditorStyle(this.contentEditor)
        this.setEditorStyle(this.configEditor)
    }
    resizeEditor() {
        const width = window.innerWidth ||
            document.documentElement.clientWidth ||
            document.body.clientWidth
        if (width > 750) {
            this.contentEditor.renderer.setPadding((width - 750) / 2)
            this.configEditor.renderer.setPadding((width - 750) / 2)
        }
        this.setEditorStyle(this.contentEditor)
        this.setEditorStyle(this.configEditor)
    }
    componentDidUpdate(prevProps, prevState) {
        if (this.props.editor.get('id') != this.props.params.id) {
            this.props.listAction.openArticle(this.props.params.id)
        }
        if (!prevProps || this.props.editor.get('id') != prevProps.editor.get('id')) {
            this.contentEditor.setValue(this.props.editor.get('content') || '', -1)
            this.configEditor.setValue(this.props.editor.get('config') || '', -1)
            this.setState({configMode: false})
        }
    }
    setEditorStyle(editor) {
        const editorOption = {
            scrollPastEnd: true,
            showGutter: false,
            wrap: true,
            theme: 'ace/theme/tomorrow',
            mode: 'ace/mode/markdown',
            showPrintMargin: false,
            fontSize: '16px',
            fontFamily: "Menlo, Consolas, 'source-code-pro', 'DejaVu Sans Mono', Monaco, 'Ubuntu Mono', 'Courier New', Courier, 'Microsoft Yahei', 'Hiragino Sans GB', 'WenQuanYi Micro Hei', monospace",
            hScrollBarAlwaysVisible: false,
            selectionStyle: "line",
            displayIndentGuides: false,
            // animatedScroll: true
        }
        editor.setOptions(editorOption)
        editor.renderer.setScrollMargin(200, 200)
        editor.container.style.lineHeight = 1.6
        editor.$blockScrolling = Infinity
        editor.on('focus', () =>
            this.props.listAction.hideList()
        )
    }
    cumulativeOffset(element) {
        let top = 0, left = 0
        do {
            top += element.offsetTop  || 0
            left += element.offsetLeft || 0
            element = element.offsetParent
        } while(element)
        return {
            top,
            left
        }
    }
    mergeState(object) {
        this.setState(Object.assign({}, this.state, object))
    }
    changeToolbar() {
        const contentSelection = this.contentEditor.getSelection()
        const cursor = contentSelection.getCursor()
        const lines = document.querySelectorAll('#content-editor .ace_line_group')
        const selectRange = this.contentEditor.getSelectionRange()
        const firstShowRow = this.contentEditor.getFirstVisibleRow()
        const startRow = selectRange.start.row - firstShowRow
        const endRow = selectRange.end.row - firstShowRow;
        ([]).forEach.call(lines, (line, idx) => {
            if (idx >= startRow - 1 && idx <= endRow + 1) {
                line.className = 'ace_line_group no-blur'
            } else {
                line.className = 'ace_line_group'
            }
        })
        _.delay(() => {
            const contentSelection = this.contentEditor.getSelection()
            const contentSession = this.contentEditor.getSession()
            const toolbarElem = document.querySelector('#editor-toolbar')
            const cursor = contentSelection.getCursor()
            const selectedText = contentSession.doc.getTextRange(contentSelection.getRange())
            const cursorElem = document.querySelector('#content-editor .ace_cursor')
            const cursorPos = this.cumulativeOffset(cursorElem)
            // this.contentEditor.scrollToLine(cursor.row, true, true)
            if (_.trim(selectedText) && !selectedText.includes('\n')) {
                toolbarElem.style.top = cursorPos.top - 45 + 'px'
                toolbarElem.style.left = cursorPos.left - 15 + 'px'
                this.mergeState({
                    toolbar: {show: true, selectMode: true}
                })
                return
            } else {
                toolbarElem.style.top = cursorPos.top - 7 + 'px'
                toolbarElem.style.left = cursorPos.left - 55 + 'px'
                this.mergeState({
                    toolbar: {selectMode: false}
                })
            }
            const line = contentSession.getLine(cursor.row)
            if (cursor.column == 0 && line.length == 0) {
                this.mergeState({
                    toolbar: {show: true}
                })
            } else {
                this.mergeState({
                    toolbar: {show: false}
                })
            }
        }, 100)
    }
    componentDidMount() {
        // init content editor
        this.contentEditor = ace.edit('content-editor')
        window.editor = this.contentEditor
        this.setEditorStyle(this.contentEditor)
        // init config editor
        this.configEditor = ace.edit('config-editor')
        this.setEditorStyle(this.configEditor)
        this.configEditor.on('input', () => {
            this.props.editorAction.setHeader(this.configEditor.getValue())
            this.onEditorChange()
        })
        this.contentEditor.on('input', () => {
            this.onEditorChange()
        })
        const contentSelection = this.contentEditor.getSelection()
        contentSelection.on('changeSelection', this.changeToolbar.bind(this))
        contentSelection.on('changeCursor', this.changeToolbar.bind(this))
        // resize by window size
        this.resizeEditor()
        window.addEventListener('resize', this.resizeEditor.bind(this))
        this.componentDidUpdate()
    }
    onEditorChange() {
        const config = _.trim(this.configEditor.getValue())
        const content = _.trim(this.contentEditor.getValue())
        const current = `${_.trim(config)}\n\n---\n\n${_.trim(content)}`
        this.props.editorAction.setCurrent(current)
    }
    moveContentEditorCursor(row, column) {
        const contentSelection = this.contentEditor.getSelection()
        const cursor = contentSelection.getCursor()
        contentSelection.moveTo(cursor.row + row, cursor.column + column)
        _.delay(() => {
            this.mergeState({
                toolbar: {show: false}
            })
        }, 200)
        this.contentEditor.focus()
    }
    onFileSelect(event) {
        this.props.editorAction.uploadImage(event.currentTarget.files[0], (path) => {
            this.contentEditor.insert(`![](${path})`)
            this.moveContentEditorCursor(0, 0)
        })
        event.currentTarget.value = ''
    }
    onActionClick(event) {
        const type = event.currentTarget.getAttribute('type')
        const contentSelection = this.contentEditor.getSelection()
        const contentSession = this.contentEditor.getSession()
        const selectedText = this.contentEditor.getSelectedText()
        switch(type) {
            case 'code':
                this.contentEditor.insert('``` \n```')
                this.moveContentEditorCursor(-1, 1)
                return
            case 'ol':
                this.contentEditor.insert('1. ')
                this.moveContentEditorCursor(0, 0)
                return
            case 'ul':
                this.contentEditor.insert('- ')
                this.moveContentEditorCursor(0, 0)
                return
            case 'indent':
                this.contentEditor.insert('> ')
                this.moveContentEditorCursor(0, 0)
                return
            case 'line':
                this.contentEditor.insert('----------\n')
                this.moveContentEditorCursor(0, 0)
                return
            case 'bold':
                {
                    this.contentEditor.insert(`**${selectedText}**`)
                    this.moveContentEditorCursor(0, -(selectedText.length + 2))
                    const cursor = contentSelection.getCursor()
                    contentSelection.selectTo(cursor.row, cursor.column + selectedText.length)
                    return
                }

            case 'italic':
                {
                    this.contentEditor.insert(`*${selectedText}*`)
                    this.moveContentEditorCursor(0, -(selectedText.length + 1))
                    const cursor = contentSelection.getCursor()
                    contentSelection.selectTo(cursor.row, cursor.column + selectedText.length)
                    return
                }
            case 'link':
                {
                    this.contentEditor.insert(`[${selectedText}]()`)
                    this.moveContentEditorCursor(0, -1)
                    return
                }
            case 'head':
                {
                    this.contentEditor.insert(`### `)
                    this.moveContentEditorCursor(0, 0)
                    return
                }
        }

    }
    render() {
        const editor = this.props.editor
        return (
            <div className="editor-wrap">
                <Header title={editor.get('title')} tags={editor.get('tags')} edit={this.state.configMode} onClick={() => this.switchConfigMode()} />
                <ul id="editor-toolbar" className={classNames({hide: !this.state.toolbar.show, select: this.state.toolbar.selectMode})}>
                    <li className="show"><i className={classNames('fa', {'fa-plus': !this.state.toolbar.selectMode}, {'fa-align-right': this.state.toolbar.selectMode})}></i></li>
                    {
                        !this.state.toolbar.selectMode ?
                        [<li className="action image">
                            <input type="file" accept="image/*" onChange={this.onFileSelect.bind(this)}/>
                            <i className="fa fa-picture-o"></i>
                            <i className="fa fa-refresh fa-spin hide"></i>
                        </li>,
                        <li type="code" className="action" onClick={this.onActionClick.bind(this)}><i className="fa fa-code"></i></li>,
                        <div className="extend">
                            <li type="line" className="action" onClick={this.onActionClick.bind(this)}><i className="fa fa-ellipsis-h"></i></li>
                            <li type="indent" className="action subaction" onClick={this.onActionClick.bind(this)}><i className="fa fa-indent"></i></li>
                            <li type="ol" className="action subaction" onClick={this.onActionClick.bind(this)}><i className="fa fa-list-ol"></i></li>
                            <li type="ul" className="action subaction" onClick={this.onActionClick.bind(this)}><i className="fa fa-list-ul"></i></li>
                            <li type="head" className="action subaction" onClick={this.onActionClick.bind(this)}><i className="fa fa-text-height"></i></li>
                        </div>] :
                        [<li type="bold" className="action" onClick={this.onActionClick.bind(this)}><i className="fa fa-bold"></i></li>,
                        <li type="italic" className="action" onClick={this.onActionClick.bind(this)}><i className="fa fa-italic"></i></li>,
                        <li type="link" className="action" onClick={this.onActionClick.bind(this)}><i className="fa fa-link"></i></li>
                        ]
                    }
                </ul>
                <div className={classNames({hide: this.state.configMode})}><div id="content-editor"></div></div>
                <div className={classNames({hide: !this.state.configMode})}><div id="config-editor"></div></div>
            </div>
        )
    }
}

export default connect(function(state) {
    return {
        editor: state.editor
    }
}, function(dispatch) {
    return {
        listAction: bindActionCreators(listAction, dispatch),
        editorAction: bindActionCreators(editorAction, dispatch),
        utilAction: bindActionCreators(utilAction, dispatch)
    }
})(Editor)
