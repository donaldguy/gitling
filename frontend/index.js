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
//cm.setValue(localStorage.getItem('gitling'));

fetch("/api/repos/gitling/README.md").then(function(response) {
  if (response.status !== 200) {
    console.log("Error fetching body: "+ response.statusText);
  } else {
    response.text().then(function(md_body){
      document.getElementById('fileName').innerText = "gitling/README.md"
      cm.setValue(md_body);
    })
  }
});

fetch("/api/repos/gitling/README.md/version").then(function (response) {
  if (response.status !== 200) {
    console.log("Error fetch version: " + response.statusText);
  } else {
    response.json().then(function (version) {
      document.getElementById('fileVersion').innerText = version.ID + " (" + version.Timestamp + ")"
    })
  }
});

// TODO: Add some debounce on this?
cm.on("change", function(_cm, _change) {
  localStorage.setItem('gitling', cm.getValue());
})

// XXX: for exposing / testing
// window.cm = cm;
// window.HyperMD = HyperMD;