# Contributing to Godocx

We welcome bug reports, feature requests, and pull requests.
Please follow these guidelines to keep things organized.

## Reporting Issues
- Search existing issues before opening a new one.
- Provide a clear description of the problem.
- Include Go version, OS, and relevant code snippets.

## Requesting Features
- Open a GitHub issue labeled `feature`.
- Describe the use case and, if possible, share a sample DOCX.

## OOXML Standards Reference

Godocx operates on DOCX files that conform to the **Office Open XML (OOXML)** standard.

**All contributors should consult these resources** when implementing or reviewing features, to ensure the generated XML is valid and compatible.

- **Official Standard:**  
  [ECMA-376 Office Open XML File Formats](https://www.ecma-international.org/publications-and-standards/standards/ecma-376/)  
  The authoritative specification for OOXML, including packaging, markup, and compatibility details.

- **Quick Schema Lookup:**  
  [Datypic OOXML Schema Reference](https://www.datypic.com/sc/ooxml/)  
  An informal, browsable view of OOXML `.xsd` schemas and elements, useful for quick tag and attribute lookups.

- **Developer Tool:**  
  [OOXML Validator for VS Code](https://marketplace.visualstudio.com/items?itemName=mikeebowen.ooxml-validator-vscode)  
  A VS Code extension that validates OOXML documents against the schema, helping catch errors early during development.
