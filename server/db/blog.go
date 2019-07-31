package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

/*
enums for `DocType` values in `Blog`
*/
var (
	DocTypeMD   int8 = 0
	DocTypeHTML int8 = 1
)

/*
Blog struct
*/
type Blog struct {
	ID            bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`
	URLIDs        string        `bson:"url_ids" json:"url_ids"`
	Title         string        `bson:"title" json:"title"`
	Description   string        `bson:"description" json:"description"`
	Author        string        `bson:"author" json:"author"`
	FormattedDate string        `bson:"formatted_date" json:"formatted_date"`
	DocType       int8          `bson:"doc_type" json:"doc_type"`
	MDSrc         string        `bson:"md_src" json:"md_src"`
	HTMLSrc       string        `bson:"html_src" json:"html_src"`
	Thumbnail     string        `bson:"thumbnail" json:"thumbnail"`
	CreatedAt     int32         `bson:"created_at" json:"created_at"`
	Likes         int           `bson:"likes" json:"likes"`
	IsFeatured    bool          `bson:"is_featured" json:"is_featured"`
	IsPublic      bool          `bson:"is_public" json:"is_public"`
	IsDeleted     bool          `bson:"is_deleted" json:"is_deleted"`
	IsSeries      bool          `bson:"is_series" json:"is_series"`
	SubBlogs      []SubBlog     `bson:"sub_blogs" json:"sub_blogs"`
}

// SubBlog - Blog model type
type SubBlog struct {
	Title         string `bson:"title" json:"title"`
	Description   string `bson:"description" json:"description"`
	FormattedDate string `bson:"formatted_date" json:"formatted_date"`
	DocType       int8   `bson:"doc_type" json:"doc_type"`
	MDSrc         string `bson:"md_src" json:"md_src"`
	HTMLSrc       string `bson:"html_src" json:"html_src"`
	Likes         int    `bson:"likes" json:"likes"`
}

// Blogs - Slice of Blogs
type Blogs []Blog

// Read - Reads all Documents from blogs
func (bs *Blogs) Read(f, s bson.M) error {
	err := MgoDB.C("blogs").Find(f).Sort("-created_at").Select(s).All(bs)
	if err != nil {
		return err
	}

	return nil
}

// Create - Creates new Document
func (b *Blog) Create() error {
	err := MgoDB.C("blogs").Insert(&b)
	if err != nil {
		return err
	}

	return nil
}

// Read - Reads single Document
func (b *Blog) Read(f, s bson.M) error {
	err := MgoDB.C("blogs").Find(f).Select(s).One(b)
	if err != nil {
		return err
	}

	return nil
}

// Update - Updates a Document by ID
func (b *Blog) Update(s bson.M, u bson.M) error {
	change := mgo.Change{
		Update:    bson.M{"$set": u},
		ReturnNew: true,
	}
	_, err := MgoDB.C("blogs").Find(s).Apply(change, b)

	return err
}

// Delete - Deletes a Document
func (b *Blog) Delete(id bson.ObjectId) error {
	// err := blogCol.Update(bson.M{"_id": id}, bson.M{"is_deleted": true})
	change := mgo.Change{
		Update:    bson.M{"$set": bson.M{"is_deleted": true}},
		ReturnNew: true,
	}
	_, err := MgoDB.C("blogs").Find(bson.M{"_id": b.ID}).Apply(change, b)

	return err
}

// DeletePermanent - Deletes a document permanently
func (b *Blog) DeletePermanent() error {
	err := MgoDB.C("blogs").Remove(bson.M{"_id": b.ID})
	return err
}