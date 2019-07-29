// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "rankdb": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/Vivino/rankdb/api/design
// --out=$(GOPATH)/src/github.com/Vivino/rankdb/api
// --version=v1.3.1

package app

import (
	"encoding/json"
	"time"
	"unicode/utf8"

	"github.com/goadesign/goa"
)

// Backup Information (default view)
//
// Identifier: application/vnd.rankdb.backup_status+json; view=default
type RankdbBackupStatus struct {
	// Will be true if backup was cancelled
	Cancelled bool `form:"cancelled" json:"cancelled" yaml:"cancelled" xml:"cancelled"`
	// Custom information provided by backup
	Custom map[string]string `form:"custom,omitempty" json:"custom,omitempty" yaml:"custom,omitempty" xml:"custom,omitempty"`
	// Will be true when the backup has finished processing
	Done bool `form:"done" json:"done" yaml:"done" xml:"done"`
	// Failed operations, indexed by list IDs
	Errors map[string]string `form:"errors,omitempty" json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty"`
	// Time backup was finished
	Finished *time.Time `form:"finished,omitempty" json:"finished,omitempty" yaml:"finished,omitempty" xml:"finished,omitempty"`
	// Number of lists to be backed up
	Lists int `form:"lists" json:"lists" yaml:"lists" xml:"lists"`
	// Number of lists backed up now
	ListsDone int `form:"lists_done" json:"lists_done" yaml:"lists_done" xml:"lists_done"`
	// Size of stored data
	Size int64 `form:"size" json:"size" yaml:"size" xml:"size"`
	// Time backup was started
	Started time.Time `form:"started" json:"started" yaml:"started" xml:"started"`
	// Storage used for backup
	Storage string `form:"storage" json:"storage" yaml:"storage" xml:"storage"`
	// URI of backed up content. Used for restore.
	URI string `form:"uri" json:"uri" yaml:"uri" xml:"uri"`
}

// Validate validates the RankdbBackupStatus media type instance.
func (mt *RankdbBackupStatus) Validate() (err error) {

	if mt.Storage == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "storage"))
	}

	if mt.URI == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uri"))
	}

	return
}

// Backup Information (full view)
//
// Identifier: application/vnd.rankdb.backup_status+json; view=full
type RankdbBackupStatusFull struct {
	// Will be true if backup was cancelled
	Cancelled bool `form:"cancelled" json:"cancelled" yaml:"cancelled" xml:"cancelled"`
	// Custom information provided by backup
	Custom map[string]string `form:"custom,omitempty" json:"custom,omitempty" yaml:"custom,omitempty" xml:"custom,omitempty"`
	// Will be true when the backup has finished processing
	Done bool `form:"done" json:"done" yaml:"done" xml:"done"`
	// Failed operations, indexed by list IDs
	Errors map[string]string `form:"errors,omitempty" json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty"`
	// Time backup was finished
	Finished *time.Time `form:"finished,omitempty" json:"finished,omitempty" yaml:"finished,omitempty" xml:"finished,omitempty"`
	// Number of lists to be backed up
	Lists int `form:"lists" json:"lists" yaml:"lists" xml:"lists"`
	// Number of lists backed up now
	ListsDone int `form:"lists_done" json:"lists_done" yaml:"lists_done" xml:"lists_done"`
	// Size of stored data
	Size int64 `form:"size" json:"size" yaml:"size" xml:"size"`
	// Time backup was started
	Started time.Time `form:"started" json:"started" yaml:"started" xml:"started"`
	// Storage used for backup
	Storage string `form:"storage" json:"storage" yaml:"storage" xml:"storage"`
	// Successful operations, list IDs
	Success []string `form:"success,omitempty" json:"success,omitempty" yaml:"success,omitempty" xml:"success,omitempty"`
	// URI of backed up content. Used for restore.
	URI string `form:"uri" json:"uri" yaml:"uri" xml:"uri"`
}

// Validate validates the RankdbBackupStatusFull media type instance.
func (mt *RankdbBackupStatusFull) Validate() (err error) {

	if mt.Storage == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "storage"))
	}

	if mt.URI == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uri"))
	}

	return
}

// Backup Information (tiny view)
//
// Identifier: application/vnd.rankdb.backup_status+json; view=tiny
type RankdbBackupStatusTiny struct {
	// Will be true if backup was cancelled
	Cancelled bool `form:"cancelled" json:"cancelled" yaml:"cancelled" xml:"cancelled"`
	// Will be true when the backup has finished processing
	Done bool `form:"done" json:"done" yaml:"done" xml:"done"`
	// Time backup was finished
	Finished *time.Time `form:"finished,omitempty" json:"finished,omitempty" yaml:"finished,omitempty" xml:"finished,omitempty"`
	// Number of lists to be backed up
	Lists int `form:"lists" json:"lists" yaml:"lists" xml:"lists"`
	// Number of lists backed up now
	ListsDone int `form:"lists_done" json:"lists_done" yaml:"lists_done" xml:"lists_done"`
	// Size of stored data
	Size int64 `form:"size" json:"size" yaml:"size" xml:"size"`
	// Time backup was started
	Started time.Time `form:"started" json:"started" yaml:"started" xml:"started"`
	// URI of backed up content. Used for restore.
	URI string `form:"uri" json:"uri" yaml:"uri" xml:"uri"`
}

// Validate validates the RankdbBackupStatusTiny media type instance.
func (mt *RankdbBackupStatusTiny) Validate() (err error) {

	if mt.URI == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "uri"))
	}

	return
}

// Backup Information (default view)
//
// Identifier: application/vnd.rankdb.callback+json; view=default
type RankdbCallback struct {
	CallbackURL string `form:"callback_url" json:"callback_url" yaml:"callback_url" xml:"callback_url"`
	ID          string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the RankdbCallback media type instance.
func (mt *RankdbCallback) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.CallbackURL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "callback_url"))
	}
	return
}

// List Element (default view)
//
// Identifier: application/vnd.rankdb.element+json; view=default
type RankdbElement struct {
	// Element rank in list from bottom.
	//  Bottom element has value 0.
	FromBottom int `form:"from_bottom" json:"from_bottom" yaml:"from_bottom" xml:"from_bottom"`
	// Element rank in list from top.
	//  Top element has value 0.
	FromTop int `form:"from_top" json:"from_top" yaml:"from_top" xml:"from_top"`
	// ID of element
	ID uint64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ListID string `form:"list_id" json:"list_id" yaml:"list_id" xml:"list_id"`
	// Custom payload. Stored untouched. On updates null means do not update. `{}` is the empty value.
	Payload json.RawMessage `form:"payload" json:"payload" yaml:"payload" xml:"payload"`
	// Score of the element. Higher score gives higher placement.
	Score uint64 `form:"score" json:"score" yaml:"score" xml:"score"`
	// Tie breaker, used if score matches for consistent sorting. Higher value = higher placement if score is equal.
	TieBreaker uint32 `form:"tie_breaker" json:"tie_breaker" yaml:"tie_breaker" xml:"tie_breaker"`
	// Date of last update
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the RankdbElement media type instance.
func (mt *RankdbElement) Validate() (err error) {

	if mt.ListID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "list_id"))
	}
	if mt.Payload == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "payload"))
	}

	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ListID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.list_id`, mt.ListID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ListID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 2, true))
	}
	if utf8.RuneCountInString(mt.ListID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 100, false))
	}
	if mt.Score < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.score`, mt.Score, 0, true))
	}
	if mt.TieBreaker < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 0, true))
	}
	if mt.TieBreaker > 4294967295 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 4294967295, false))
	}
	return
}

// List Element (full view)
//
// Identifier: application/vnd.rankdb.element+json; view=full
type RankdbElementFull struct {
	// Element rank in list from bottom.
	//  Bottom element has value 0.
	FromBottom int `form:"from_bottom" json:"from_bottom" yaml:"from_bottom" xml:"from_bottom"`
	// Element rank in list from top.
	//  Top element has value 0.
	FromTop int `form:"from_top" json:"from_top" yaml:"from_top" xml:"from_top"`
	// ID of element
	ID uint64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ListID string `form:"list_id" json:"list_id" yaml:"list_id" xml:"list_id"`
	// Local element rank in list from bottom when requesting sub-list.
	//  Bottom element has value 0.
	LocalFromBottom *int `form:"local_from_bottom,omitempty" json:"local_from_bottom,omitempty" yaml:"local_from_bottom,omitempty" xml:"local_from_bottom,omitempty"`
	// Local element rank in list from top when requesting sub-list.
	//  Top element has value 0.
	LocalFromTop *int `form:"local_from_top,omitempty" json:"local_from_top,omitempty" yaml:"local_from_top,omitempty" xml:"local_from_top,omitempty"`
	// Neighbors to current element
	Neighbors *struct {
		// Elements above the current element.
		// Ends with element just above current.
		Above RankdbElementCollection `form:"above,omitempty" json:"above,omitempty" yaml:"above,omitempty" xml:"above,omitempty"`
		// Elements below the current element.
		// Starts with element just below current.
		Below RankdbElementCollection `form:"below,omitempty" json:"below,omitempty" yaml:"below,omitempty" xml:"below,omitempty"`
	} `form:"neighbors,omitempty" json:"neighbors,omitempty" yaml:"neighbors,omitempty" xml:"neighbors,omitempty"`
	// Custom payload. Stored untouched. On updates null means do not update. `{}` is the empty value.
	Payload json.RawMessage `form:"payload" json:"payload" yaml:"payload" xml:"payload"`
	// Score of the element. Higher score gives higher placement.
	Score uint64 `form:"score" json:"score" yaml:"score" xml:"score"`
	// Tie breaker, used if score matches for consistent sorting. Higher value = higher placement if score is equal.
	TieBreaker uint32 `form:"tie_breaker" json:"tie_breaker" yaml:"tie_breaker" xml:"tie_breaker"`
	// Date of last update
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the RankdbElementFull media type instance.
func (mt *RankdbElementFull) Validate() (err error) {

	if mt.ListID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "list_id"))
	}
	if mt.Payload == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "payload"))
	}

	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ListID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.list_id`, mt.ListID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ListID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 2, true))
	}
	if utf8.RuneCountInString(mt.ListID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 100, false))
	}
	if mt.Neighbors != nil {
		if err2 := mt.Neighbors.Above.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := mt.Neighbors.Below.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Score < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.score`, mt.Score, 0, true))
	}
	if mt.TieBreaker < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 0, true))
	}
	if mt.TieBreaker > 4294967295 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 4294967295, false))
	}
	return
}

// List Element (full-update view)
//
// Identifier: application/vnd.rankdb.element+json; view=full-update
type RankdbElementFullUpdate struct {
	// Element rank in list from bottom.
	//  Bottom element has value 0.
	FromBottom int `form:"from_bottom" json:"from_bottom" yaml:"from_bottom" xml:"from_bottom"`
	// Element rank in list from top.
	//  Top element has value 0.
	FromTop int `form:"from_top" json:"from_top" yaml:"from_top" xml:"from_top"`
	// ID of element
	ID uint64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ListID string `form:"list_id" json:"list_id" yaml:"list_id" xml:"list_id"`
	// Local element rank in list from bottom when requesting sub-list.
	//  Bottom element has value 0.
	LocalFromBottom *int `form:"local_from_bottom,omitempty" json:"local_from_bottom,omitempty" yaml:"local_from_bottom,omitempty" xml:"local_from_bottom,omitempty"`
	// Local element rank in list from top when requesting sub-list.
	//  Top element has value 0.
	LocalFromTop *int `form:"local_from_top,omitempty" json:"local_from_top,omitempty" yaml:"local_from_top,omitempty" xml:"local_from_top,omitempty"`
	// Neighbors to current element
	Neighbors *struct {
		// Elements above the current element.
		// Ends with element just above current.
		Above RankdbElementCollection `form:"above,omitempty" json:"above,omitempty" yaml:"above,omitempty" xml:"above,omitempty"`
		// Elements below the current element.
		// Starts with element just below current.
		Below RankdbElementCollection `form:"below,omitempty" json:"below,omitempty" yaml:"below,omitempty" xml:"below,omitempty"`
	} `form:"neighbors,omitempty" json:"neighbors,omitempty" yaml:"neighbors,omitempty" xml:"neighbors,omitempty"`
	// Custom payload. Stored untouched. On updates null means do not update. `{}` is the empty value.
	Payload json.RawMessage `form:"payload" json:"payload" yaml:"payload" xml:"payload"`
	// Rank of element before update
	PreviousRank *RankdbElement `form:"previous_rank,omitempty" json:"previous_rank,omitempty" yaml:"previous_rank,omitempty" xml:"previous_rank,omitempty"`
	// Score of the element. Higher score gives higher placement.
	Score uint64 `form:"score" json:"score" yaml:"score" xml:"score"`
	// Tie breaker, used if score matches for consistent sorting. Higher value = higher placement if score is equal.
	TieBreaker uint32 `form:"tie_breaker" json:"tie_breaker" yaml:"tie_breaker" xml:"tie_breaker"`
	// Date of last update
	UpdatedAt time.Time `form:"updated_at" json:"updated_at" yaml:"updated_at" xml:"updated_at"`
}

// Validate validates the RankdbElementFullUpdate media type instance.
func (mt *RankdbElementFullUpdate) Validate() (err error) {

	if mt.ListID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "list_id"))
	}
	if mt.Payload == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "payload"))
	}

	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ListID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.list_id`, mt.ListID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ListID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 2, true))
	}
	if utf8.RuneCountInString(mt.ListID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.list_id`, mt.ListID, utf8.RuneCountInString(mt.ListID), 100, false))
	}
	if mt.Neighbors != nil {
		if err2 := mt.Neighbors.Above.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
		if err2 := mt.Neighbors.Below.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.PreviousRank != nil {
		if err2 := mt.PreviousRank.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Score < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.score`, mt.Score, 0, true))
	}
	if mt.TieBreaker < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 0, true))
	}
	if mt.TieBreaker > 4294967295 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.tie_breaker`, mt.TieBreaker, 4294967295, false))
	}
	return
}

// List Element (tiny view)
//
// Identifier: application/vnd.rankdb.element+json; view=tiny
type RankdbElementTiny struct {
	// Element rank in list from bottom.
	//  Bottom element has value 0.
	FromBottom int `form:"from_bottom" json:"from_bottom" yaml:"from_bottom" xml:"from_bottom"`
	// Element rank in list from top.
	//  Top element has value 0.
	FromTop int `form:"from_top" json:"from_top" yaml:"from_top" xml:"from_top"`
	// ID of element
	ID uint64 `form:"id" json:"id" yaml:"id" xml:"id"`
	// Custom payload. Stored untouched. On updates null means do not update. `{}` is the empty value.
	Payload json.RawMessage `form:"payload" json:"payload" yaml:"payload" xml:"payload"`
	// Score of the element. Higher score gives higher placement.
	Score uint64 `form:"score" json:"score" yaml:"score" xml:"score"`
}

// Validate validates the RankdbElementTiny media type instance.
func (mt *RankdbElementTiny) Validate() (err error) {

	if mt.Payload == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "payload"))
	}
	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
	}
	if mt.Score < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.score`, mt.Score, 0, true))
	}
	return
}

// RankdbElementCollection is the media type for an array of RankdbElement (default view)
//
// Identifier: application/vnd.rankdb.element+json; type=collection; view=default
type RankdbElementCollection []*RankdbElement

// Validate validates the RankdbElementCollection media type instance.
func (mt RankdbElementCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RankdbElementCollection is the media type for an array of RankdbElement (full view)
//
// Identifier: application/vnd.rankdb.element+json; type=collection; view=full
type RankdbElementFullCollection []*RankdbElementFull

// Validate validates the RankdbElementFullCollection media type instance.
func (mt RankdbElementFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RankdbElementCollection is the media type for an array of RankdbElement (full-update view)
//
// Identifier: application/vnd.rankdb.element+json; type=collection; view=full-update
type RankdbElementFullUpdateCollection []*RankdbElementFullUpdate

// Validate validates the RankdbElementFullUpdateCollection media type instance.
func (mt RankdbElementFullUpdateCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RankdbElementCollection is the media type for an array of RankdbElement (tiny view)
//
// Identifier: application/vnd.rankdb.element+json; type=collection; view=tiny
type RankdbElementTinyCollection []*RankdbElementTiny

// Validate validates the RankdbElementTinyCollection media type instance.
func (mt RankdbElementTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// Result of a list operation. Will contain either an error or a list. (default view)
//
// Identifier: application/vnd.rankdb.listopresult+json; view=default
type RankdbListopresult struct {
	// Error, if any encountered
	Error *string         `form:"error,omitempty" json:"error,omitempty" yaml:"error,omitempty" xml:"error,omitempty"`
	List  *RankdbRanklist `form:"list,omitempty" json:"list,omitempty" yaml:"list,omitempty" xml:"list,omitempty"`
}

// Validate validates the RankdbListopresult media type instance.
func (mt *RankdbListopresult) Validate() (err error) {
	if mt.List != nil {
		if err2 := mt.List.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// RankdbListsresult media type (default view)
//
// Identifier: application/vnd.rankdb.listsresult+json; view=default
type RankdbListsresult struct {
	Lists RankdbRanklistCollection `form:"lists" json:"lists" yaml:"lists" xml:"lists"`
	// The number of lists after the last element
	ListsAfter int `form:"lists_after" json:"lists_after" yaml:"lists_after" xml:"lists_after"`
	// The number of lists before the first element
	ListsBefore int `form:"lists_before" json:"lists_before" yaml:"lists_before" xml:"lists_before"`
}

// Validate validates the RankdbListsresult media type instance.
func (mt *RankdbListsresult) Validate() (err error) {

	if mt.Lists == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "lists"))
	}
	if err2 := mt.Lists.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// RankdbMultielement media type (default view)
//
// Identifier: application/vnd.rankdb.multielement+json; view=default
type RankdbMultielement struct {
	// Elements that was found in the list, ordered by score.
	Found RankdbElementCollection `form:"found" json:"found" yaml:"found" xml:"found"`
	// Elements that wasn't found in the list. Unordered.
	NotFound []uint64 `form:"not_found,omitempty" json:"not_found,omitempty" yaml:"not_found,omitempty" xml:"not_found,omitempty"`
}

// Validate validates the RankdbMultielement media type instance.
func (mt *RankdbMultielement) Validate() (err error) {
	if mt.Found == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "found"))
	}
	if err2 := mt.Found.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// RankdbOperation_success media type (default view)
//
// Identifier: application/vnd.rankdb.operation_success+json; view=default
type RankdbOperationSuccess struct {
	// If `results` parameter was true, the resulting element is returned here.
	Results RankdbElementTinyCollection `form:"results,omitempty" json:"results,omitempty" yaml:"results,omitempty" xml:"results,omitempty"`
}

// Validate validates the RankdbOperationSuccess media type instance.
func (mt *RankdbOperationSuccess) Validate() (err error) {
	if err2 := mt.Results.Validate(); err2 != nil {
		err = goa.MergeErrors(err, err2)
	}
	return
}

// Rank List (default view)
//
// Identifier: application/vnd.rankdb.ranklist+json; view=default
type RankdbRanklist struct {
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Load Index on server startup
	LoadIndex bool `form:"load_index" json:"load_index" yaml:"load_index" xml:"load_index"`
	// Merge adjacent Segments with less than this number of entries
	MergeSize int `form:"merge_size" json:"merge_size" yaml:"merge_size" xml:"merge_size"`
	// Custom metadata. String to String hash.
	Metadata map[string]string `form:"metadata" json:"metadata" yaml:"metadata" xml:"metadata"`
	// Set used for storage
	Set string `form:"set" json:"set" yaml:"set" xml:"set"`
	// Split Segments larger than this number of entries
	SplitSize int `form:"split_size" json:"split_size" yaml:"split_size" xml:"split_size"`
}

// Validate validates the RankdbRanklist media type instance.
func (mt *RankdbRanklist) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Set == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "set"))
	}
	if mt.Metadata == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "metadata"))
	}

	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 2, true))
	}
	if utf8.RuneCountInString(mt.ID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 100, false))
	}
	if mt.MergeSize < 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.merge_size`, mt.MergeSize, 10, true))
	}
	if mt.MergeSize > 100000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.merge_size`, mt.MergeSize, 100000, false))
	}
	if utf8.RuneCountInString(mt.Set) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.set`, mt.Set, utf8.RuneCountInString(mt.Set), 2, true))
	}
	if utf8.RuneCountInString(mt.Set) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.set`, mt.Set, utf8.RuneCountInString(mt.Set), 100, false))
	}
	if mt.SplitSize < 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.split_size`, mt.SplitSize, 10, true))
	}
	if mt.SplitSize > 100000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.split_size`, mt.SplitSize, 100000, false))
	}
	return
}

// Rank List (full view)
//
// Identifier: application/vnd.rankdb.ranklist+json; view=full
type RankdbRanklistFull struct {
	// Average number of elements per segment
	AvgSegmentElements float64        `form:"avg_segment_elements" json:"avg_segment_elements" yaml:"avg_segment_elements" xml:"avg_segment_elements"`
	BottomElement      *RankdbElement `form:"bottom_element,omitempty" json:"bottom_element,omitempty" yaml:"bottom_element,omitempty" xml:"bottom_element,omitempty"`
	// Cache hits while segments have been loaded.
	CacheHits int `form:"cache_hits" json:"cache_hits" yaml:"cache_hits" xml:"cache_hits"`
	// Cache misses while segments have been loaded.
	CacheMisses int `form:"cache_misses" json:"cache_misses" yaml:"cache_misses" xml:"cache_misses"`
	// Cache hit percentage while segments have been loaded.
	CachePercent float64 `form:"cache_percent" json:"cache_percent" yaml:"cache_percent" xml:"cache_percent"`
	// Number of elements in list
	Elements int `form:"elements" json:"elements" yaml:"elements" xml:"elements"`
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Load Index on server startup
	LoadIndex bool `form:"load_index" json:"load_index" yaml:"load_index" xml:"load_index"`
	// Merge adjacent Segments with less than this number of entries
	MergeSize int `form:"merge_size" json:"merge_size" yaml:"merge_size" xml:"merge_size"`
	// Custom metadata. String to String hash.
	Metadata map[string]string `form:"metadata" json:"metadata" yaml:"metadata" xml:"metadata"`
	// Number of segment in list
	Segments int `form:"segments" json:"segments" yaml:"segments" xml:"segments"`
	// Set used for storage
	Set string `form:"set" json:"set" yaml:"set" xml:"set"`
	// Split Segments larger than this number of entries
	SplitSize int `form:"split_size" json:"split_size" yaml:"split_size" xml:"split_size"`
	// This element is only returned if 'top_bottom' parameter is set to true.
	TopElement *RankdbElement `form:"top_element,omitempty" json:"top_element,omitempty" yaml:"top_element,omitempty" xml:"top_element,omitempty"`
}

// Validate validates the RankdbRanklistFull media type instance.
func (mt *RankdbRanklistFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Set == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "set"))
	}
	if mt.Metadata == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "metadata"))
	}

	if mt.BottomElement != nil {
		if err2 := mt.BottomElement.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 2, true))
	}
	if utf8.RuneCountInString(mt.ID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 100, false))
	}
	if mt.MergeSize < 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.merge_size`, mt.MergeSize, 10, true))
	}
	if mt.MergeSize > 100000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.merge_size`, mt.MergeSize, 100000, false))
	}
	if utf8.RuneCountInString(mt.Set) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.set`, mt.Set, utf8.RuneCountInString(mt.Set), 2, true))
	}
	if utf8.RuneCountInString(mt.Set) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.set`, mt.Set, utf8.RuneCountInString(mt.Set), 100, false))
	}
	if mt.SplitSize < 10 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.split_size`, mt.SplitSize, 10, true))
	}
	if mt.SplitSize > 100000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.split_size`, mt.SplitSize, 100000, false))
	}
	if mt.TopElement != nil {
		if err2 := mt.TopElement.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Rank List (tiny view)
//
// Identifier: application/vnd.rankdb.ranklist+json; view=tiny
type RankdbRanklistTiny struct {
	// The ID of the list to apply the operation on.
	// Can be `a` to `z` (both upper/lower case), `0` to `9` or one of these characters `_-.`
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the RankdbRanklistTiny media type instance.
func (mt *RankdbRanklistTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9-_.]+$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[a-zA-Z0-9-_.]+$`))
	}
	if utf8.RuneCountInString(mt.ID) < 2 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 2, true))
	}
	if utf8.RuneCountInString(mt.ID) > 100 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 100, false))
	}
	return
}

// RankdbRanklistCollection is the media type for an array of RankdbRanklist (default view)
//
// Identifier: application/vnd.rankdb.ranklist+json; type=collection; view=default
type RankdbRanklistCollection []*RankdbRanklist

// Validate validates the RankdbRanklistCollection media type instance.
func (mt RankdbRanklistCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RankdbRanklistCollection is the media type for an array of RankdbRanklist (tiny view)
//
// Identifier: application/vnd.rankdb.ranklist+json; type=collection; view=tiny
type RankdbRanklistTinyCollection []*RankdbRanklistTiny

// Validate validates the RankdbRanklistTinyCollection media type instance.
func (mt RankdbRanklistTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// RankdbRestoreresult media type (default view)
//
// Identifier: application/vnd.rankdb.restoreresult+json; view=default
type RankdbRestoreresult struct {
	// Failed operations, indexed by list IDs
	Errors map[string]string `form:"errors" json:"errors" yaml:"errors" xml:"errors"`
	// Successful restore operations, indexed by list IDs
	Restored int `form:"restored" json:"restored" yaml:"restored" xml:"restored"`
	// Skipped lists, indexed by list IDs
	Skipped int `form:"skipped" json:"skipped" yaml:"skipped" xml:"skipped"`
}

// Validate validates the RankdbRestoreresult media type instance.
func (mt *RankdbRestoreresult) Validate() (err error) {

	if mt.Errors == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "errors"))
	}
	return
}

// RankdbResultlist media type (default view)
//
// Identifier: application/vnd.rankdb.resultlist+json; view=default
type RankdbResultlist struct {
	// Failed operations, indexed by list IDs
	Errors map[string]string `form:"errors,omitempty" json:"errors,omitempty" yaml:"errors,omitempty" xml:"errors,omitempty"`
	// Successful operations, indexed by list IDs
	Success map[string]*RankdbOperationSuccess `form:"success,omitempty" json:"success,omitempty" yaml:"success,omitempty" xml:"success,omitempty"`
}

// Validate validates the RankdbResultlist media type instance.
func (mt *RankdbResultlist) Validate() (err error) {
	for _, e := range mt.Success {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// System Info. The model is sparse and may contain other information. (default view)
//
// Identifier: application/vnd.rankdb.sysinfo+json; view=default
type RankdbSysinfo struct {
	// Element Cache Information
	ElementCache map[string]interface{} `form:"element_cache,omitempty" json:"element_cache,omitempty" yaml:"element_cache,omitempty" xml:"element_cache,omitempty"`
	// Lazy saver cache information.
	LazySaver json.RawMessage `form:"lazy_saver,omitempty" json:"lazy_saver,omitempty" yaml:"lazy_saver,omitempty" xml:"lazy_saver,omitempty"`
	// Memory Information
	Memory map[string]interface{} `form:"memory,omitempty" json:"memory,omitempty" yaml:"memory,omitempty" xml:"memory,omitempty"`
}