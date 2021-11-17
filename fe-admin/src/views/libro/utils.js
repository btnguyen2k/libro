/*
Libro utilities.

@author Thanh Nguyen <btnguyen2k@gmail.com>
@libro
*/

// import hljs from 'highlight.js'          // all languages
import hljs from 'highlight.js/lib/common'  // only common languages
// import hljs from 'highlight.js/lib/core' // no language
import 'highlight.js/styles/default.css'
import katex from 'katex'
import 'katex/contrib/mhchem/mhchem'
import 'katex/dist/katex.min.css'
import {marked} from "marked"
import DOMPurify from "dompurify"

function htmlDecode(html) {
    const doc = new DOMParser().parseFromString(html, "text/html")
    return doc.documentElement.textContent
}

//ref: https://github.com/markedjs/marked/issues/1538#issuecomment-575838181
let katexId = 0
const nextKatexId = () => `__special_katext_id_${katexId++}__`
const mathExpMap = {}
const myRenderer = new marked.Renderer()

function replaceMathWithIds(text, el) {
    // allowing newlines inside of `$$...$$`
    text = text.replace(/\$\$([\s\S]+?)\$\$/g, (_match, expression) => {
        const id = nextKatexId()
        mathExpMap[id] = {type: 'block', expression, el: el}
        return id
    })

    // Not allowing newlines inside of `$...$`
    text = text.replace(/\$([^\n]+?)\$/g, (_match, expression) => {
        const id = nextKatexId()
        mathExpMap[id] = {type: 'inline', expression, el: el}
        return id
    })

    return text
}

const orgListitem = myRenderer.listitem
myRenderer.listitem = function (text, task, checked) {
    return orgListitem(replaceMathWithIds(htmlDecode(text), 'listitem'), task, checked)
}

const orgParagraph = myRenderer.paragraph
myRenderer.paragraph = function (text) {
    return orgParagraph(replaceMathWithIds(htmlDecode(text), 'paragraph'))
}

const orgTablecell = myRenderer.tablecell
myRenderer.tablecell = function (content, flags) {
    return orgTablecell(replaceMathWithIds(htmlDecode(content), 'tablecell'), flags)
}

const orgText = myRenderer.text
myRenderer.text = function (text) {
    return orgText(replaceMathWithIds(htmlDecode(text), 'text'))
}

marked.use({
    gfm: true,
    renderer: myRenderer,
    langPrefix: 'hljs language-', // highlight.js css expects a top-level 'hljs' class
    highlight: function (code, lang) {
        const language = hljs.getLanguage(lang) ? lang : 'plaintext'
        return hljs.highlight(code, { language }).value
    },
})

function markdownRender(markdownInput, sanitize) {
    const html = marked.parse(markdownInput)
    const latexHtml = html.replace(/(__special_katext_id_\d+__)/g, (_match, capture) => {
        const token = mathExpMap[capture]
        return katex.renderToString(token.expression, {
            displayMode: token.type == 'block',
            output: 'html',
            throwOnError: false
        })
    })
    return sanitize ? latexHtml : DOMPurify.sanitize(latexHtml, {ADD_ATTR: ['target']})
}

export {
    markdownRender
}
