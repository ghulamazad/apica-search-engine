import React, { useRef, useState } from "react";
import axios, { AxiosError } from "axios";

const UploadButton: React.FC = () => {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [uploading, setUploading] = useState(false);
  const [message, setMessage] = useState("");

  const handleUpload = async () => {
    if (fileInputRef.current?.files?.length) {
      const formData = new FormData();
      formData.append("files", fileInputRef.current.files[0]);

      try {
        setUploading(true);
        const res = await axios.post("http://localhost:8080/upload", formData, {
          headers: { "Content-Type": "multipart/form-data" },
        });
        setMessage("File uploaded successfully");
      } catch (err) {
        console.error("Upload error:", err);
        if (err instanceof AxiosError) {
          if (
            err.response?.data?.error?.includes(
              "unsupported MIME type uploaded"
            )
          ) {
            setMessage("The uploaded file format is not supported. Please upload a valid file.");
          } else {
            setMessage("Upload failed!");
          }
        } else {
          setMessage((err as Error).message ?? "Upload failed!");
        }
      } finally {
        setUploading(false);
      }
    }
  };

  return (
    <div className="my-6 flex flex-col items-center gap-4">
      <input
        type="file"
        ref={fileInputRef}
        className="hidden"
        onChange={handleUpload}
      />
      <button
        onClick={() => fileInputRef.current?.click()}
        className="bg-green-600 hover:bg-green-700 active:bg-green-800 text-white font-semibold py-3 px-6 rounded-md shadow-md transition duration-300 ease-in-out"
      >
        {uploading ? "Uploading..." : "Upload Parquet File"}
      </button>
      {message && (
        <p className="text-center text-sm text-gray-600">{message}</p>
      )}
    </div>
  );
};

export default UploadButton;
