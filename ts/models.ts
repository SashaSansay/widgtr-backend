/* Do not change, this code is generated from Golang structs */


export enum WidgetType {
    PARAGRAPH = "paragraph",
    HEADING_H1 = "heading_h1",
    HEADING_H2 = "heading_h2",
    HEADING_H3 = "heading_h3",
    HEADING_H4 = "heading_h4",
    HEADING_H5 = "heading_h5",
    HEADING_H6 = "heading_h6",
}
export interface User {
    ID: number;
    FirstName: string;
    LastName: string;
}
export interface HeadingH6Content {
    text: string;
    color: string;
}
export interface HeadingH5Content {
    text: string;
    color: string;
}
export interface HeadingH4Content {
    text: string;
    color: string;
}
export interface HeadingH3Content {
    text: string;
    color: string;
}
export interface HeadingH2Content {
    text: string;
    color: string;
}
export interface HeadingH1Content {
    text: string;
    color: string;
}
export interface ParagraphContent {
    text: string;
    color: string;
}
export interface Widget {
    id: number[];
    type: WidgetType;
    children: Widget[];
    paragraph?: ParagraphContent;
    heading_h1?: HeadingH1Content;
    heading_h2?: HeadingH2Content;
    heading_h3?: HeadingH3Content;
    heading_h4?: HeadingH4Content;
    heading_h5?: HeadingH5Content;
    heading_h6?: HeadingH6Content;
}
export interface TextContent {
    text: string;
    color: string;
}






