# SourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SyntaxVersion** | Pointer to **string** | The syntax version of the source, e.g. &#x60;cel1&#x60;. | [optional] 
**Location** | Pointer to **string** | The location name. All position information attached to an expression is relative to this location.  The location could be a file, UI element, or similar. For example, &#x60;acme/app/AnvilPolicy.cel&#x60;. | [optional] 
**LineOffsets** | Pointer to **[]int32** | Monotonically increasing list of code point offsets where newlines &#x60;\\n&#x60; appear.  The line number of a given position is the index &#x60;i&#x60; where for a given &#x60;id&#x60; the &#x60;line_offsets[i] &lt; id_positions[id] &lt; line_offsets[i+1]&#x60;. The column may be derivd from &#x60;id_positions[id] - line_offsets[i]&#x60;. | [optional] 
**Positions** | Pointer to **map[string]int32** | A map from the parse node id (e.g. &#x60;Expr.id&#x60;) to the code point offset within the source. | [optional] 
**MacroCalls** | Pointer to [**map[string]Expr**](Expr.md) | A map from the parse node id where a macro replacement was made to the call &#x60;Expr&#x60; that resulted in a macro expansion.  For example, &#x60;has(value.field)&#x60; is a function call that is replaced by a &#x60;test_only&#x60; field selection in the AST. Likewise, the call &#x60;list.exists(e, e &gt; 10)&#x60; translates to a comprehension expression. The key in the map corresponds to the expression id of the expanded macro, and the value is the call &#x60;Expr&#x60; that was replaced. | [optional] 

## Methods

### NewSourceInfo

`func NewSourceInfo() *SourceInfo`

NewSourceInfo instantiates a new SourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSourceInfoWithDefaults

`func NewSourceInfoWithDefaults() *SourceInfo`

NewSourceInfoWithDefaults instantiates a new SourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSyntaxVersion

`func (o *SourceInfo) GetSyntaxVersion() string`

GetSyntaxVersion returns the SyntaxVersion field if non-nil, zero value otherwise.

### GetSyntaxVersionOk

`func (o *SourceInfo) GetSyntaxVersionOk() (*string, bool)`

GetSyntaxVersionOk returns a tuple with the SyntaxVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyntaxVersion

`func (o *SourceInfo) SetSyntaxVersion(v string)`

SetSyntaxVersion sets SyntaxVersion field to given value.

### HasSyntaxVersion

`func (o *SourceInfo) HasSyntaxVersion() bool`

HasSyntaxVersion returns a boolean if a field has been set.

### GetLocation

`func (o *SourceInfo) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SourceInfo) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SourceInfo) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SourceInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetLineOffsets

`func (o *SourceInfo) GetLineOffsets() []int32`

GetLineOffsets returns the LineOffsets field if non-nil, zero value otherwise.

### GetLineOffsetsOk

`func (o *SourceInfo) GetLineOffsetsOk() (*[]int32, bool)`

GetLineOffsetsOk returns a tuple with the LineOffsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineOffsets

`func (o *SourceInfo) SetLineOffsets(v []int32)`

SetLineOffsets sets LineOffsets field to given value.

### HasLineOffsets

`func (o *SourceInfo) HasLineOffsets() bool`

HasLineOffsets returns a boolean if a field has been set.

### GetPositions

`func (o *SourceInfo) GetPositions() map[string]int32`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *SourceInfo) GetPositionsOk() (*map[string]int32, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *SourceInfo) SetPositions(v map[string]int32)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *SourceInfo) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetMacroCalls

`func (o *SourceInfo) GetMacroCalls() map[string]Expr`

GetMacroCalls returns the MacroCalls field if non-nil, zero value otherwise.

### GetMacroCallsOk

`func (o *SourceInfo) GetMacroCallsOk() (*map[string]Expr, bool)`

GetMacroCallsOk returns a tuple with the MacroCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacroCalls

`func (o *SourceInfo) SetMacroCalls(v map[string]Expr)`

SetMacroCalls sets MacroCalls field to given value.

### HasMacroCalls

`func (o *SourceInfo) HasMacroCalls() bool`

HasMacroCalls returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


