# APIView token parsers

This page describes how to contribute to [APIView](../../../src//dotnet/APIView/APIViewWeb/APIViewWeb.csproj) language level parsers.
Specifically how to create or update a language parser to produce a hierarchy of the API using tokens with context, instead of a flat token list.

## Benefits of token schema with parent - child context

- Ability to granularly identify a specific node and it's sub nodes. For e.g. if a review line for a class is available then all it's child nodes contain all methods and properties within that class.
- Faster diffing using tree shaker instead of current text based comparison.
- Provide diffing with context of where the line that changed belong to in the tree, instead of showing the 5 lines before and after the change.
- Provide cross language view for a granular section within an API review.
- Less data to be stored in token file which are located in azure storage blob.

## Key concepts

APIView token schema is available in [TypeSpec](./apiview-treestyle-parser-schema/model.tsp) and [JSON](./apiview-treestyle-parser-schema/model.json). Language parsers needs to create a `CodeFile` object as per the schema.
`CodeFile` object contains metadata about the package and an array of `ReviewLine` objects. Each object of `ReviewLine` in `CodeFile` object is a top-level line to be listed. For e.g. Top-level lines are mostly namespace or module level nodes.
Each `ReviewLine` object has children of `ReviewLine` to include sub nodes that needs to be listed in a review.

A sample token file is present [here](./apiview-treestyle-parser-schema/sample/Azure.Template_token.json).

APIView generates a navigation tree based on the information in the tokens. A token is included in the navigation tree if `NavigatonDisplayName` is set in `ReviewToken` and `LineId` is set in `ReviewLine` object that contains the `ReviewToken`


## Serialization

Serialize the generated code file to JSON. The output file should have `.json` extension. Try to make the json as small as possible by ignoring null values and empty collections.
Don't worry about indentation as it will be handled by the tree structure based on the parent-child relationship among `ReviewLine` objects.

## Examples

A sample token file for .NET package `Azure.Template` is present [here](./apiview-treestyle-parser-schema/sample/Azure.Template_token.json) and corresponding rendered text representation is available [here](./apiview-treestyle-parser-schema/sample/Azure.Template_review_content.txt).


## JSON token validation

You can validate JSON tokens against required JSON schema using [JSON schema validator](https://www.jsonschemavalidator.net/).

- Select `Custom` as schema type and copy and paste the contents in [json schema](./apiview-treestyle-parser-schema/model.json) to left panel.
- Generate APIView token file and paste generated JSON content onto right side panel to validate.


## Get help

Please reach out at [APIView Teams Channel](https://teams.microsoft.com/l/channel/19%3A3adeba4aa1164f1c889e148b1b3e3ddd%40thread.skype/APIView?groupId=3e17dcb0-4257-4a30-b843-77f47f1d4121&tenantId=72f988bf-86f1-41af-91ab-2d7cd011db47) if you need more information.
