package main

import (
	"context"
	"fmt"
	"log"

	"github.com/y-mabuchi/grpc-golang/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog Client")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Yusuke",
		Title:    "My First Blog",
		Content:  "Content of the first blog",
	}

	res, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}
	fmt.Printf("Blog has created: %v\n", res)
	blogId := res.GetBlog().GetId()

	// read blog
	fmt.Println("Reading the blog")

	_, err2 := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
		BlogId: "dateatoiukjd",
	})
	if err2 != nil {
		fmt.Printf("Error happened while readng: %v\n", err2)
	}

	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, readBlogErr := c.ReadBlog(context.Background(), readBlogReq)
	if readBlogErr != nil {
		fmt.Printf("Error happened while readng: %v\n", readBlogErr)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	// update Blog
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Changed Author",
		Title:    "My First blog (edited)",
		Content:  "Content of the first blog, with some awesome additions!",
	}
	updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
		Blog: newBlog,
	})
	if updateErr != nil {
		fmt.Printf("Error happened while update: %v\n", updateErr)
	}

	fmt.Printf("Blog was updated: %v\n", updateRes)

	// delete blog
	deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
		BlogId: blogId,
	})
	if deleteErr != nil {
		fmt.Printf("Error happened while deleting: %v\n", deleteErr)
	}

	fmt.Printf("Blog was deleted: %v\n", deleteRes)
}
