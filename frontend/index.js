const CodeMirror = require('codemirror/lib/codemirror')
const HyperMD = require('hypermd')

 // languages supported for highlighting
require('codemirror/mode/go/go') 
require('codemirror/mode/shell/shell') 

const cm = HyperMD.fromTextArea(document.getElementById("content"), {
  mode: {
    name: 'hypermd'
  }
})
cm.setSize(null, "100%");
cm.setValue(localStorage.getItem('gitling'));

// TODO: Add some debounce on this?
cm.on("change", function(_cm, _change) {
  localStorage.setItem('gitling', cm.getValue());
})

// XXX: for exposing / testing
// window.cm = cm;
// window.HyperMD = HyperMD;