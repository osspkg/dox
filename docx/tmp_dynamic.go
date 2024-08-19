package docx

var tmpDynamicWordDocument = dynamicFile{
	Filepath: "/word/document.xml",
	AfterBody: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:o="urn:schemas-microsoft-com:office:office"
    xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships"
    xmlns:v="urn:schemas-microsoft-com:vml"
    xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"
    xmlns:w10="urn:schemas-microsoft-com:office:word"
    xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing"
    xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape"
    xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup"
    xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006"
    xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing"
    xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml"
    xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" mc:Ignorable="w14 wp14 w15">
    <w:body>`,
	BeforeBody: `<w:sectPr>
            <w:headerReference w:type="default" r:id="rId2" />
            <w:footerReference w:type="default" r:id="rId3" />
            <w:type w:val="nextPage" />
            <w:pgSz w:w="11906" w:h="16838" />
            <w:pgMar w:left="1134" w:right="1134" w:gutter="0" w:header="1134" w:top="1693"
                w:footer="1134" w:bottom="1693" />
            <w:pgNumType w:fmt="decimal" />
            <w:formProt w:val="false" />
            <w:textDirection w:val="lrTb" />
        </w:sectPr>
	</w:body>
</w:document>`,
}

var tmpDynamicWordStyles = dynamicFile{
	Filepath: "/word/styles.xml",
	AfterBody: `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
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
    <w:style w:type="paragraph" w:styleId="Normal">
        <w:name w:val="Normal" />
        <w:qFormat />
        <w:pPr>
            <w:widowControl />
            <w:bidi w:val="0" />
        </w:pPr>
        <w:rPr>
            <w:rFonts w:ascii="Liberation Serif" w:hAnsi="Liberation Serif"
                w:eastAsia="Noto Serif CJK SC" w:cs="Lohit Devanagari" />
            <w:color w:val="auto" />
            <w:kern w:val="2" />
            <w:sz w:val="24" />
            <w:szCs w:val="24" />
            <w:lang w:val="en-US" w:eastAsia="zh-CN" w:bidi="hi-IN" />
        </w:rPr>
    </w:style>`,
	BeforeBody: `</w:styles>`,
}
