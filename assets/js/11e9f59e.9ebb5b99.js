"use strict";(self.webpackChunkgodocx_docs=self.webpackChunkgodocx_docs||[]).push([[402],{8204:(e,n,d)=>{d.r(n),d.d(n,{assets:()=>c,contentTitle:()=>a,default:()=>h,frontMatter:()=>i,metadata:()=>r,toc:()=>o});var l=d(4848),t=d(8453);const i={},a="Tables",r={id:"basics/tables",title:"Tables",description:"This section provides an overview of how to create and manipulate tables using the godocx library. We'll cover the basics of adding a table to a document, styling it, and populating it with data.",source:"@site/docs/basics/tables.md",sourceDirName:"basics",slug:"/basics/tables",permalink:"/godocx/docs/basics/tables",draft:!1,unlisted:!1,editUrl:"https://github.com/facebook/docusaurus/tree/main/packages/create-docusaurus/templates/shared/docs/basics/tables.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Paragraphs",permalink:"/godocx/docs/basics/paragraphs"}},c={},o=[{value:"Adding a Table",id:"adding-a-table",level:2},{value:"Example: Adding and Styling a Table",id:"example-adding-and-styling-a-table",level:3},{value:"Another Example",id:"another-example",level:3},{value:"Table Styles",id:"table-styles",level:3},{value:"Further Reading",id:"further-reading",level:2}];function s(e){const n={a:"a",code:"code",h1:"h1",h2:"h2",h3:"h3",li:"li",p:"p",pre:"pre",ul:"ul",...(0,t.R)(),...e.components};return(0,l.jsxs)(l.Fragment,{children:[(0,l.jsx)(n.h1,{id:"tables",children:"Tables"}),"\n",(0,l.jsxs)(n.p,{children:["This section provides an overview of how to create and manipulate tables using the ",(0,l.jsx)(n.code,{children:"godocx"})," library. We'll cover the basics of adding a table to a document, styling it, and populating it with data."]}),"\n",(0,l.jsx)(n.h2,{id:"adding-a-table",children:"Adding a Table"}),"\n",(0,l.jsxs)(n.p,{children:["To add a table to a document, use the ",(0,l.jsx)(n.code,{children:"AddTable"})," method. You can then style the table and add rows and cells as needed."]}),"\n",(0,l.jsx)(n.h3,{id:"example-adding-and-styling-a-table",children:"Example: Adding and Styling a Table"}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-go",children:'package main\n\nimport (\n    "log"\n    "github.com/gomutex/godocx"\n)\n\nfunc main() {\n    document, err := godocx.NewDocument()\n    if err != nil {\n        log.Fatal(err)\n    }\n\n   records := []struct{ Qty, ID, Desc string }{\n        {"5", "A001", "Laptop"},\n        {"10", "B202", "Smartphone"},\n        {"2", "E505", "Smartwatch"},\n    }\n\n    // Add a new table to the document\n\ttable := document.AddTable()\n\n    // Set the table style \n\ttable.Style("LightList-Accent4")\n\n    // Add New row to the table\n\thdrRow := table.AddRow()\n\n    // Add cells and add parargraph to those cells\n\thdrRow.AddCell().AddParagraph("Qty")\n\thdrRow.AddCell().AddParagraph("ID")\n\thdrRow.AddCell().AddParagraph("Description")\n\n\tfor _, record := range records {\n\t\trow := table.AddRow()\n\t\trow.AddCell().AddParagraph(record.Qty)\n\t\trow.AddCell().AddParagraph(record.ID)\n\t\trow.AddCell().AddParagraph(record.Desc)\n\t}\n\n    // Save the document\n    err = document.SaveTo("table.docx")\n    if err != nil {\n        log.Fatal(err)\n    }\n}\n'})}),"\n",(0,l.jsx)(n.p,{children:"In this example:"}),"\n",(0,l.jsxs)(n.ul,{children:["\n",(0,l.jsx)(n.li,{children:"A new table is added to the document."}),"\n",(0,l.jsx)(n.li,{children:"The table is styled using a predefined style."}),"\n",(0,l.jsx)(n.li,{children:'A header row is added with three columns: "Qty", "ID", and "Description".'}),"\n",(0,l.jsx)(n.li,{children:"Data rows are added to the table from a slice of records."}),"\n"]}),"\n",(0,l.jsx)(n.h3,{id:"another-example",children:"Another Example"}),"\n",(0,l.jsx)(n.p,{children:"Here is one more example, showing how to create a table, add rows and cells, and save the document."}),"\n",(0,l.jsx)(n.pre,{children:(0,l.jsx)(n.code,{className:"language-go",children:'package main\n\nimport (\n    "log"\n    "github.com/gomutex/godocx"\n)\n\nfunc main() {\n    document, err := godocx.NewDocument()\n    if err != nil {\n        log.Fatal(err)\n    }\n\n    // Add a new table to the document\n    table := document.AddTable()\n    // Apply a predefined style to the table\n    table.Style("LightList-Accent2")\n\n    // Add the first row\n    tblRow := table.AddRow()\n    cell00 := tblRow.AddCell()\n    cell00.AddParagraph("Column1")\n    cell01 := tblRow.AddCell()\n    cell01.AddParagraph("Column2")\n\n    // Add the second row\n    tblRow1 := table.AddRow()\n    cell10 := tblRow1.AddCell()\n    cell10.AddParagraph("Row2 - Column 1")\n    cell11 := tblRow1.AddCell()\n    cell11.AddParagraph("Row2 - Column 2")\n\n    // Save the document\n    err = document.SaveTo("table.docx")\n    if err != nil {\n        log.Fatal(err)\n    }\n}\n'})}),"\n",(0,l.jsx)(n.p,{children:"In this example:"}),"\n",(0,l.jsxs)(n.ul,{children:["\n",(0,l.jsx)(n.li,{children:"A new table is added and styled."}),"\n",(0,l.jsx)(n.li,{children:"Two rows are added, each with two cells containing text."}),"\n",(0,l.jsx)(n.li,{children:"The document is saved to a file named table.docx."}),"\n"]}),"\n",(0,l.jsx)(n.h3,{id:"table-styles",children:"Table Styles"}),"\n",(0,l.jsx)(n.p,{children:"The following table styles can be used in the default template:"}),"\n",(0,l.jsxs)(n.ul,{children:["\n",(0,l.jsx)(n.li,{children:'"LightShading"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent1"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent2"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent3"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent4"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent5"'}),"\n",(0,l.jsx)(n.li,{children:'"LightShading-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"LightList"'}),"\n",(0,l.jsx)(n.li,{children:'"LightList-Accent1" ... "LightList-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"LightGrid"'}),"\n",(0,l.jsx)(n.li,{children:'"LightGrid-Accent1" ... "LightGrid-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumShading"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumShading-Accent1" ... "MediumShading-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumShading2"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumShading2-Accent1" ... "MediumShading2-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumList1"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumList1-Accent1" ... "MediumList1-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumList2"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumList2-Accent1" ... "MediumList2-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"TableGrid"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid1"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid1-Accent1" ... "MediumGrid1-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid2"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid2-Accent1" ... "MediumGrid2-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid3"'}),"\n",(0,l.jsx)(n.li,{children:'"MediumGrid3-Accent1" ... "MediumGrid3-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"DarkList"'}),"\n",(0,l.jsx)(n.li,{children:'"DarkList-Accent1" ... "DarkList-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulShading"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulShading-Accent1" ... "ColorfulShading-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulList"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulList-Accent1" ... "ColorfulList-Accent6"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulGrid"'}),"\n",(0,l.jsx)(n.li,{children:'"ColorfulGrid-Accent1" ... "ColorfulGrid-Accent6"'}),"\n"]}),"\n",(0,l.jsx)(n.h2,{id:"further-reading",children:"Further Reading"}),"\n",(0,l.jsxs)(n.p,{children:["For more information on the full range of table-related functions available in the ",(0,l.jsx)(n.code,{children:"godocx"})," library, please refer to the ",(0,l.jsx)(n.a,{href:"https://pkg.go.dev/github.com/gomutex/godocx/docx#Table",children:"official documentation on pkg.go.dev"}),"."]})]})}function h(e={}){const{wrapper:n}={...(0,t.R)(),...e.components};return n?(0,l.jsx)(n,{...e,children:(0,l.jsx)(s,{...e})}):s(e)}},8453:(e,n,d)=>{d.d(n,{R:()=>a,x:()=>r});var l=d(6540);const t={},i=l.createContext(t);function a(e){const n=l.useContext(i);return l.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function r(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:a(e.components),l.createElement(i.Provider,{value:n},e.children)}}}]);