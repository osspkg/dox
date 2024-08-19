package docx

var tmplContentTypes = staticFile{
	Filepath: "[Content_Types].xml",
	Body: `<?xml version="1.0" encoding="UTF-8"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
    <Default Extension="xml" ContentType="application/xml" />
    <Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml" />
    <Default Extension="png" ContentType="image/png" />
    <Default Extension="jpeg" ContentType="image/jpeg" />
    <Override PartName="/_rels/.rels"
        ContentType="application/vnd.openxmlformats-package.relationships+xml" />
    <Override PartName="/docProps/core.xml"
        ContentType="application/vnd.openxmlformats-package.core-properties+xml" />
    <Override PartName="/docProps/app.xml"
        ContentType="application/vnd.openxmlformats-officedocument.extended-properties+xml" />
    <Override PartName="/word/_rels/document.xml.rels"
        ContentType="application/vnd.openxmlformats-package.relationships+xml" />
    <Override PartName="/word/document.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" />
    <Override PartName="/word/styles.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.styles+xml" />
    <Override PartName="/word/header1.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml" />
    <Override PartName="/word/footer1.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml" />
    <Override PartName="/word/settings.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.settings+xml" />
    <Override PartName="/word/numbering.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.numbering+xml" />
    <Override PartName="/word/fontTable.xml"
        ContentType="application/vnd.openxmlformats-officedocument.wordprocessingml.fontTable+xml" />
</Types>
`,
}

var tmplDocPropsApp = staticFile{
	Filepath: "/docProps/app.xml",
	Body: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Properties xmlns="http://schemas.openxmlformats.org/officeDocument/2006/extended-properties"
    xmlns:vt="http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes">
    <Template></Template>
    <TotalTime>0</TotalTime>
    <Application>go-dox/docx</Application>
    <AppVersion>1.0000</AppVersion>
    <Pages>1</Pages>
    <Words>26</Words>
    <Characters>148</Characters>
    <CharactersWithSpaces>148</CharactersWithSpaces>
    <Paragraphs>24</Paragraphs>
</Properties>
`,
}

var tmplDocPropsCore = staticFile{
	Filepath: "/docProps/core.xml",
	Body: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<cp:coreProperties
    xmlns:cp="http://schemas.openxmlformats.org/package/2006/metadata/core-properties"
    xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:dcterms="http://purl.org/dc/terms/"
    xmlns:dcmitype="http://purl.org/dc/dcmitype/"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
    <dcterms:created xsi:type="dcterms:W3CDTF">{{date .DateTime}}</dcterms:created>
    <dc:creator></dc:creator>
    <dc:description></dc:description>
    <dc:language>{{.Lang}}</dc:language>
    <cp:lastModifiedBy></cp:lastModifiedBy>
    <dcterms:modified xsi:type="dcterms:W3CDTF">{{date .DateTime}}</dcterms:modified>
    <cp:revision>1</cp:revision>
    <dc:subject></dc:subject>
    <dc:title></dc:title>
</cp:coreProperties>
`,
}

var tmplRels = staticFile{
	Filepath: "/_rels/.rels",
	Body: `<?xml version="1.0" encoding="UTF-8"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
    <Relationship Id="rId1"
        Type="http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
        Target="docProps/core.xml" />
    <Relationship Id="rId2"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
        Target="docProps/app.xml" />
    <Relationship Id="rId3"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
        Target="word/document.xml" />
</Relationships>
`,
}

var tmplWordDocumentRels = staticFile{
	Filepath: "/word/_rels/document.xml.rels",
	Body: `<?xml version="1.0" encoding="UTF-8"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
    <Relationship Id="rId1"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
        Target="styles.xml" />
    <Relationship Id="rId2"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"
        Target="header1.xml" />
    <Relationship Id="rId3"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"
        Target="footer1.xml" />
    <Relationship Id="rId4"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"
        Target="numbering.xml" />
    <Relationship Id="rId5"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"
        Target="fontTable.xml" />
    <Relationship Id="rId6"
        Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"
        Target="settings.xml" />
</Relationships>
`,
}

var tmplWordFontTable = staticFile{
	Filepath: "/word/fontTable.xml",
	Body: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:fonts xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
    xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
    <w:font w:name="Times New Roman">
        <w:charset w:val="00" />
        <w:family w:val="roman" />
        <w:pitch w:val="variable" />
    </w:font>
    <w:font w:name="Symbol">
        <w:charset w:val="02" />
        <w:family w:val="roman" />
        <w:pitch w:val="variable" />
    </w:font>
    <w:font w:name="Arial">
        <w:charset w:val="00" />
        <w:family w:val="swiss" />
        <w:pitch w:val="variable" />
    </w:font>
    <w:font w:name="Liberation Serif">
        <w:altName w:val="Times New Roman" />
        <w:charset w:val="01" />
        <w:family w:val="roman" />
        <w:pitch w:val="variable" />
    </w:font>
    <w:font w:name="Liberation Sans">
        <w:altName w:val="Arial" />
        <w:charset w:val="01" />
        <w:family w:val="swiss" />
        <w:pitch w:val="variable" />
    </w:font>
</w:fonts>
`,
}

var tmplDefaultWordStyles = staticFile{
	Filepath: "/word/styles.xml",
	Body: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:styles xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
    xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" mc:Ignorable="w14">
    <w:docDefaults>
        <w:rPrDefault>
            <w:rPr>
                <w:rFonts w:ascii="Liberation Serif" w:hAnsi="Liberation Serif"
                    w:eastAsia="Noto Serif CJK SC" w:cs="Lohit Devanagari" />
                <w:kern w:val="2" />
                <w:sz w:val="24" />
                <w:szCs w:val="24" />
                <w:lang w:val="ru-RU" w:eastAsia="zh-CN" w:bidi="hi-IN" />
            </w:rPr>
        </w:rPrDefault>
        <w:pPrDefault>
            <w:pPr>
                <w:widowControl />
                <w:suppressAutoHyphens w:val="true" />
            </w:pPr>
        </w:pPrDefault>
    </w:docDefaults>
	{{.Styles}}
</w:styles>
`,
}
